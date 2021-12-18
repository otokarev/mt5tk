package symbol

import (
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	symbol2 "github.com/otokarev/mt5tk/pkg/model/entities"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var symbol string
	var group string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get symbol's details. With --group parameter it returns symbol details for a specified group.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var results symbol2.Symbol

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
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	getCmd.MarkFlagRequired("symbol")
	getCmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")

	return getCmd
}
