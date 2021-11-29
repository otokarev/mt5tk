package group

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var group string
	var jsonPath string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get group's details",
		Run: func(cmd *cobra.Command, args []string) {
			group, _ = cmd.Flags().GetString("group")
			results := modelFactory.Group().Get(group)
			jsonPath, _ := cmd.Flags().GetString("jsonpath")
			err := util.PrintResult(jsonPath, results)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&jsonPath, "jsonpath", "j", "", "JSONPath template")
	getCmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	getCmd.MarkFlagRequired("group")

	return getCmd
}
