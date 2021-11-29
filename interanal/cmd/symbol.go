package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var symbolCmd = &cobra.Command{
	Use:   "symbol",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("symbol called")
	},
}

func init() {
	rootCmd.AddCommand(symbolCmd)
}
