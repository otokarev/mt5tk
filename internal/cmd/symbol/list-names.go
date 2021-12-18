package symbol

import (
	"github.com/otokarev/mt5tk/internal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildListNames() *cobra.Command {
	var outputFormat string
	listCmd := &cobra.Command{
		Use:   "list-names",
		Short: "List symbols' names",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := modelFactory.Symbol().ListNames()
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
