package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var symbol string

var symbolGetCmd = &cobra.Command{
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

func init() {
	symbolCmd.AddCommand(symbolGetCmd)
	symbolGetCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol title")
	symbolGetCmd.MarkFlagRequired("symbol")
}
