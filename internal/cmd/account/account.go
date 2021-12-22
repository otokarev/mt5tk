package account

import (
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
)

var modelFactory *model.Factory

func Build(factory *model.Factory) *cobra.Command {
	modelFactory = factory
	var cmd = &cobra.Command{
		Use:   "account",
		Short: "Account operations",
	}

	cmd.AddCommand(buildGet())

	return cmd
}
