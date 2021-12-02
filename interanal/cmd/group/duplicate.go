package group

import (
	"github.com/spf13/cobra"
	"log"
)

func buildDuplicate() *cobra.Command {
	var srcName string
	var newName string

	var cmd = &cobra.Command{
		Use:   "duplicate",
		Short: "Duplicate group",
		Run: func(cmd *cobra.Command, args []string) {
			src, err := modelFactory.Group().Get(srcName)
			if err != nil {
				log.Fatalf("cannot load source group %s, error: %s", srcName, err.Error())
			}
			if _, err := modelFactory.Group().Get(newName); err == nil {
				log.Fatalf("group %s already exists", newName)
			}

			src.Group = newName
			if _, err := modelFactory.Group().Add(src); err != nil {
				log.Fatalf("cannot create group %s, error: %s", newName, err.Error)
			}
		},
	}
	cmd.Flags().StringVarP(&newName, "name", "n", "", "New group name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&srcName, "group", "g", "", "Source group")
	cmd.MarkFlagRequired("source-group")

	return cmd
}
