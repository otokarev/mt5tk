package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var symbolListCmd = &cobra.Command{
	Use:   "list",
	Short: "List symbols",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(modelFactory.Symbol().List())
	},
}

func init() {
	symbolCmd.AddCommand(symbolListCmd)
}
