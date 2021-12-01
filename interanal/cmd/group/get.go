package group

import (
	"github.com/otokarev/mt5tk/interanal/cmd/util/output"
	"github.com/spf13/cobra"
	"log"
)

func buildGet() *cobra.Command {
	var group string
	var outputFormat string

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get group's details",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := modelFactory.Group().Get(group)
			if err != nil {
				log.Fatal(err)
			}

			if err := output.Print(outputFormat, results); err != nil {
				log.Fatal(err)
			}
		},
	}
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "json|jsonpath=<pattern>")
	getCmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	getCmd.MarkFlagRequired("group")

	return getCmd
}
