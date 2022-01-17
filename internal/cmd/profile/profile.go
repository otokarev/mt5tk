package profile

import (
	"github.com/go-playground/validator/v10"
	"github.com/otokarev/mt5tk/internal/configurator"
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
)

var cfg = configurator.LoadConfig("config", "yaml", ".")
var validate = validator.New()

func Build(factory *model.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "profile",
		Short: "Profile operations",
	}
	cmd.AddCommand(buildSwitch())
	cmd.AddCommand(buildShow())
	cmd.AddCommand(buildNew())
	cmd.AddCommand(buildDelete())

	return cmd
}

func buildShow() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "show current configuration settings",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
