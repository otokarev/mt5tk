package group

import (
	"encoding/json"
	group2 "github.com/otokarev/mt5tk/pkg/model/entities"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func buildCreate() *cobra.Command {
	var group string
	var templatePath string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create new group by template",
		Run: func(cmd *cobra.Command, args []string) {
			exists, err := modelFactory.Group().Exists(group)
			if err != nil {
				log.Fatalf("cannot verify group %s existance, error: %s", group, err.Error())
			}
			if exists {
				log.Fatalf("group %s already exists", group)
			}

			data, err := ioutil.ReadFile(templatePath)
			if err != nil {
				log.Fatal("cannot read template file, error: ", err.Error())
			}
			var newGroup group2.Group
			if err := json.Unmarshal(data, &newGroup); err != nil {
				log.Fatal("cannot parse template file contents, error: ", err.Error())
			}

			newGroup.Group = group

			if _, err := modelFactory.Group().Add(newGroup); err != nil {
				log.Fatalf("cannot create group %s, error: %s", group, err.Error())
			}
		},
	}
	cmd.Flags().StringVarP(
		&templatePath,
		"template-path",
		"t",
		"",
		"Path to group template file (json)",
	)
	cmd.MarkFlagRequired("template-path")
	cmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	cmd.MarkFlagRequired("group")

	return cmd
}
