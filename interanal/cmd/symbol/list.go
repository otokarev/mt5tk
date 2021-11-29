package symbol

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util"
	"github.com/spf13/cobra"
	"log"
)

func buildList() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List symbols",
		Run: func(cmd *cobra.Command, args []string) {
			results := modelFactory.Symbol().List()
			jsonPath, _ := cmd.Flags().GetString("jsonpath")
			err := util.PrintResult(jsonPath, results)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return listCmd
}
