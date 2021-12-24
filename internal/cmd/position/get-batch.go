package position

import (
	"fmt"
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func buildGetBatch() *cobra.Command {
	var user string
	var symbol string
	var position string
	var group string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get-batch",
		Short: "Get batch of positions",
		Run: func(cmd *cobra.Command, args []string) {
			positions := stringToArray(position)
			symbols := stringToArray(symbol)
			users := stringToArray(user)
			groups := stringToArray(group)
			if len(users) == 0 && len(positions) == 0 && len(groups) == 0 {
				errString := "Error: at least one of the parameters \"user\", \"group\", \"position\" must be defined"
				fmt.Println(errString)
				cmd.Help()
				fmt.Println("\n" + errString)
				os.Exit(1)
			}
			if len(symbols) > 0 && len(users) == 0 {
				log.Fatal("symbol parameter does not work without user")
			}
			results, err := modelFactory.Position().GetBatch(positions, users, groups, symbols)
			if err != nil {
				log.Fatal(err)
			}

			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")
	getCmd.Flags().StringVarP(&user, "user", "u", "", "user identifier(s) separated by comma")
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "symbol identifier(s) separated by comma")
	getCmd.Flags().StringVarP(&group, "group", "g", "", "group identifier(s) separated by comma")
	getCmd.Flags().StringVarP(&position, "position", "k", "", "position identifier(s) separated by comma")

	return getCmd
}

func stringToArray(in string) []string {
	split := strings.Split(in, ",")
	var ret []string
	for i := range split {
		trimmed := strings.TrimSpace(split[i])
		if trimmed != "" {
			ret = append(ret, trimmed)
		}
	}
	return ret
}
