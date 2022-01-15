package symbol

import (
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	"github.com/otokarev/mt5tk/pkg/model/entities"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var symbol string
	var group string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get symbol's details",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var results entities.Symbol

			if group == "" {
				results, err = modelFactory.Symbol().Get(symbol)
			} else {
				results, err = modelFactory.Symbol().GetForGroup(symbol, group)
			}

			if err != nil {
				log.Fatal(err)
			}

			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&symbol, "symbol", "k", "", "symbol name")
	getCmd.MarkFlagRequired("symbol")
	getCmd.Flags().StringVarP(&group, "group", "g", "", "group name.  If flag set, command returns group's settings for the symbol")
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")

	return getCmd
}
