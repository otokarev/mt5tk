package symbol

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var symbol string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get symbol's details",
		Run: func(cmd *cobra.Command, args []string) {
			symbol, _ = cmd.Flags().GetString("symbol")
			results := modelFactory.Symbol().Get(symbol)
			err := output.Print(outputFormat, results)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")
	getCmd.MarkFlagRequired("symbol")

	return getCmd
}
