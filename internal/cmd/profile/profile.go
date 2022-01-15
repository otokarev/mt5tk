package profile

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/otokarev/mt5tk/internal/configurator"
	model2 "github.com/otokarev/mt5tk/internal/configurator/entities"
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
	"log"
)

func Build(factory *model.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "profile",
		Short: "Profile operations",
	}
	cmd.AddCommand(buildSwitch())
	cmd.AddCommand(buildShow())
	cmd.AddCommand(buildNew())

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

func buildNew() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "create new configuration settings profile",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

func buildSwitch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "switch",
		Short: "switch default config profile",
		Run: func(cmd *cobra.Command, args []string) {
			var profile string
			cfg := configurator.LoadConfig("config", "yaml", ".")
			profiles := cfg.GetProfileNames()
			current := cfg.GetDefaultProfile()
			var opts []string
			for _, p := range profiles {
				var opt string
				if p == current {
					opt = fmt.Sprintf("%s (default)", p)
				} else {
					opt = p
				}
				opts = append(opts, opt)
			}
			opts = append(opts, "Create New Profile")
			prompt := &survey.Select{
				Message: "Choose a profile:",
				Options: opts,
			}
			err := survey.AskOne(prompt, &profile)
			if err != nil {
				log.Fatal(err)
			}
			if profile == "Create New Profile" {
				createNewProfile()
			} else {
				cfg.SwitchToProfile(profile)
				fmt.Println("selected", profile)
			}
		},
	}

	return cmd
}

func createNewProfile() {
	var profile model2.Profile

	var qs = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Enter a new profile name:"},
			Validate: survey.Required,
		},
		{
			Name:     "serverAddress",
			Prompt:   &survey.Input{Message: "Enter a server address (format: <IP address>:<port>):"},
			Validate: survey.Required,
		},
		{
			Name:     "login",
			Prompt:   &survey.Input{Message: "Enter login:"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Enter password:"},
			Validate: survey.Required,
		},
	}
	err := survey.Ask(qs, &profile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("New profile: %#v", profile)
}
