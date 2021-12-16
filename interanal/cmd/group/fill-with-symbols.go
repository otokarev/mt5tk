package group

import (
	"encoding/json"
	group2 "github.com/otokarev/mt5tk/pkg/model/group"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func buildFillWithSymbols() *cobra.Command {
	var group string
	var templatePath string

	var cmd = &cobra.Command{
		Use:   "fill-with-symbols",
		Short: "Fill the groups with all symbols available",
		Long:  "Fill the groups with all symbols available. Default settings for the symbols are taken from a template file.",
		Run: func(cmd *cobra.Command, args []string) {
			groupObject, err := modelFactory.Group().Get(group)
			if err != nil {
				log.Fatalf("cannot load group %s, error: %s", group, err.Error())
			}

			data, err := ioutil.ReadFile(templatePath)
			if err != nil {
				log.Fatal("cannot read template file, error: ", err.Error())
			}
			var tmpl group2.SymbolObject
			if err := json.Unmarshal(data, &tmpl); err != nil {
				log.Fatal("cannot parse template file contents, error: ", err.Error())
			}

			symbolObjects, err := modelFactory.Symbol().ListByNames()
			if err != nil {
				log.Fatal(err)
			}

			var groupSymbolObjects []group2.SymbolObject
			for _, symbolObject := range symbolObjects {
				obj := tmpl
				obj.Path = symbolObject.Path
				groupSymbolObjects = append(groupSymbolObjects, obj)
			}

			groupObject.Symbols = groupSymbolObjects

			if _, err := modelFactory.Group().Add(groupObject); err != nil {
				log.Fatalf("cannot update egroup %s, error: %s", group, err.Error)
			}
		},
	}
	cmd.Flags().StringVarP(
		&templatePath,
		"template-path",
		"t",
		"",
		"Path to symbol template file (json)",
	)
	cmd.MarkFlagRequired("template-path")
	cmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	cmd.MarkFlagRequired("group")

	return cmd
}
