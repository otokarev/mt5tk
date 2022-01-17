package profile

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"log"
)

func buildSwitch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "switch",
		Short: "switch default config profile",
		Run: func(cmd *cobra.Command, args []string) {
			var profile string
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
				profile = createNewProfile()
			}
			cfg.SwitchToProfile(profile)
			fmt.Println("selected", profile)
		},
	}

	return cmd
}
