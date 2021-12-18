package symbol

import (
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
)

var modelFactory *model.Factory

func Build(factory *model.Factory) *cobra.Command {
	modelFactory = factory
	var cmd = &cobra.Command{
		Use:   "symbol",
		Short: "Group's symbol operations",
	}

	cmd.AddCommand(buildAdd())
	cmd.AddCommand(buildReplace())
	cmd.AddCommand(buildRemove())

	return cmd
}
