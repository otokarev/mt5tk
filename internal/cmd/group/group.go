package group

import (
	"github.com/otokarev/mt5tk/internal/cmd/group/symbol"
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
)

var modelFactory *model.Factory

func Build(factory *model.Factory) *cobra.Command {
	modelFactory = factory
	var groupCmd = &cobra.Command{
		Use:   "group",
		Short: "Groups operations",
	}

	groupCmd.AddCommand(buildGet())
	groupCmd.AddCommand(buildList())
	groupCmd.AddCommand(buildFillWithSymbols())
	groupCmd.AddCommand(buildDuplicate())
	groupCmd.AddCommand(buildCreate())

	groupCmd.AddCommand(symbol.Build(modelFactory))

	return groupCmd
}
