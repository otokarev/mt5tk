package position

import (
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var login string
	var symbol string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get position's details",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := modelFactory.Position().Get(login, symbol)
			if err != nil {
				log.Fatal(err)
			}

			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")
	getCmd.Flags().StringVarP(&login, "user", "u", "", "user identifier")
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "symbol identifier")
	getCmd.MarkFlagRequired("user")
	getCmd.MarkFlagRequired("symbol")

	return getCmd
}
