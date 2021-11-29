package group

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util"
	"github.com/spf13/cobra"
	"log"
)

func buildList() *cobra.Command {
	var jsonPath string

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List groups",
		Run: func(cmd *cobra.Command, args []string) {
			results := modelFactory.Group().List()
			jsonPath, _ := cmd.Flags().GetString("jsonpath")
			err := util.PrintResult(jsonPath, results)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	listCmd.Flags().StringVarP(&jsonPath, "jsonpath", "j", "", "JSONPath template")

	return listCmd
}
