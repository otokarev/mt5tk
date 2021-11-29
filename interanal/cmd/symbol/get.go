package symbol

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var symbol string
	var jsonPath string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get symbol's details",
		Run: func(cmd *cobra.Command, args []string) {
			symbol, _ = cmd.Flags().GetString("symbol")
			results := modelFactory.Symbol().Get(symbol)
			jsonPath, _ := cmd.Flags().GetString("jsonpath")
			err := util.PrintResult(jsonPath, results)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	getCmd.Flags().StringVarP(&jsonPath, "jsonpath", "j", "", "JSONPath template")
	getCmd.MarkFlagRequired("symbol")

	return getCmd
}
