package symbol

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildList() *cobra.Command {
	var outputFormat string
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List symbols",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := modelFactory.Symbol().List()
			if err != nil {
				log.Fatal(err)
			}
			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}

	listCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")

	return listCmd
}
