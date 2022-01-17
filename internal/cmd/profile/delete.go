package profile

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"log"
)

func buildDelete() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete MT5 server connection profile",
		Run: func(cmd *cobra.Command, args []string) {
			var result struct {
				Profile string
				Confirm bool
			}
			profiles := cfg.GetProfileNames()
			current := cfg.GetDefaultProfile()
			var opts []string
			for _, p := range profiles {
				if p != current {
					opts = append(opts, p)
				}
			}
			if len(opts) == 0 {
				log.Fatal("no profiles to remove")
			}
			var qs = []*survey.Question{
				{
					Name: "profile",
					Prompt: &survey.Select{
						Message: "Choose a profile delete (default profile is not listed):",
						Options: opts,
					},
				},
				{
					Name:     "confirm",
					Prompt:   &survey.Confirm{Message: "Are you sure?"},
					Validate: survey.Required,
				},
			}
			err := survey.Ask(qs, &result)
			if err != nil {
				log.Fatal(err)
			}
			if result.Confirm {
				cfg.DeleteProfile(result.Profile)
				log.Printf("profile `%s` deleted\n", result.Profile)
			} else {
				log.Println("no profile deleted")
			}
		},
	}

	return cmd
}
