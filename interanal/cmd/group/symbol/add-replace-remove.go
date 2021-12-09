package symbol

import (
	"encoding/json"
	group2 "github.com/otokarev/mt5tk/pkg/model/group"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func buildAdd() *cobra.Command {
	var group string
	var symbol string
	var path string
	var templatePath string

	var cmd = &cobra.Command{
		Use:   "add",
		Short: "Add a symbol by use of template",
		Run: func(cmd *cobra.Command, args []string) {
			updateSymbolWithTemplate(&group, &symbol, &path, &templatePath, false)
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
	cmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Symbol path")

	return cmd
}

func buildReplace() *cobra.Command {
	var group string
	var symbol string
	var path string
	var templatePath string

	var cmd = &cobra.Command{
		Use:   "replace",
		Short: "Replace a symbol by use of template",
		Run: func(cmd *cobra.Command, args []string) {
			updateSymbolWithTemplate(&group, &symbol, &path, &templatePath, true)
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
	cmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Symbol path")

	return cmd
}

func buildRemove() *cobra.Command {
	var group string
	var symbol string
	var path string

	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove a symbol by its name or path",
		Run: func(cmd *cobra.Command, args []string) {
			removeSymbol(&group, &symbol, &path)
		},
	}
	cmd.Flags().StringVarP(&group, "group", "g", "", "Group name")
	cmd.MarkFlagRequired("group")
	cmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol name")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Symbol path")

	return cmd
}

func removeSymbol(group *string, symbol *string, path *string) {
	groupObject, _ := validateParamsAndPrepareGroupForUpdateWithTemplate(group, symbol, path, true)

	if _, err := modelFactory.Group().Add(groupObject); err != nil {
		log.Fatalf("cannot save modified group %s, error: %s", *group, err.Error())
	}
}

func updateSymbolWithTemplate(group *string, symbol *string, path *string, templatePath *string, replace bool) {
	symbolObject := loadSymbolTemplateFromFile(templatePath)
	groupObject, pathValue := validateParamsAndPrepareGroupForUpdateWithTemplate(group, symbol, path, replace)

	symbolObject.Path = pathValue
	groupObject.Symbols = append(groupObject.Symbols, symbolObject)

	if _, err := modelFactory.Group().Add(groupObject); err != nil {
		log.Fatalf("cannot save modified group %s, error: %s", *group, err.Error())
	}
}

func validateParamsAndPrepareGroupForUpdateWithTemplate(group *string, symbol *string, path *string, replace bool) (group2.GroupObject, string) {
	if *path == "" && *symbol == "" {
		log.Fatal("Nor symbol neither path are specified")
	}
	if *path != "" && *symbol != "" {
		log.Fatal("Either symbol or path must be specified")
	}
	exists, err := modelFactory.Group().Exists(*group)
	if err != nil {
		log.Fatalf("cannot verify group %s existance, error: %s", *group, err.Error())
	}
	if exists == false {
		log.Fatalf("group %s does not exist", *group)
	}

	if *path == "" {
		symbolObject, err := modelFactory.Symbol().Get(*symbol)
		if err != nil {
			log.Fatalf("cannot load symbol %s, error: %s", *symbol, err.Error())
		}
		*path = symbolObject.Path
	}

	groupObject, err := modelFactory.Group().Get(*group)
	if err != nil {
		log.Fatalf("cannot load group %s, error: %s", *group, err.Error())
	}

	idxToDelete := -1
	for i, s := range groupObject.Symbols {
		if s.Path == *path {
			if replace {
				idxToDelete = i
			} else {
				log.Fatalf("symbol's path %s already resides in the group %s", *path, *group)
			}
		}
	}
	if replace {
		if idxToDelete >= 0 {
			groupObject.Symbols = append(groupObject.Symbols[:idxToDelete], groupObject.Symbols[idxToDelete+1:]...)
		} else {
			log.Fatalf("symbol's path %s is not found in the group %s", *path, *group)
		}
	}

	return groupObject, *path
}

func loadSymbolTemplateFromFile(templatePath *string) group2.SymbolObject {
	data, err := ioutil.ReadFile(*templatePath)
	if err != nil {
		log.Fatal("cannot read template file, error: ", err.Error())
	}
	var s group2.SymbolObject
	if err := json.Unmarshal(data, &s); err != nil {
		log.Fatal("cannot parse template file contents, error: ", err.Error())
	}

	return s
}
