package symbol

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var symbol string
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Show symbol's details",
		Run: func(cmd *cobra.Command, args []string) {
			symbol, _ = cmd.Flags().GetString("symbol")
			json, err := json.MarshalIndent(modelFactory.Symbol().Get(symbol), "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(json))
		},
	}
	getCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol title")
	getCmd.MarkFlagRequired("symbol")

	return getCmd
}
