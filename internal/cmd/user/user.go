package user

import (
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
)

var modelFactory *model.Factory

func Build(factory *model.Factory) *cobra.Command {
	modelFactory = factory
	var symbolCmd = &cobra.Command{
		Use:   "user",
		Short: "User operations",
	}

	symbolCmd.AddCommand(buildGet())

	return symbolCmd
}
