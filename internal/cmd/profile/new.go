package profile

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	model2 "github.com/otokarev/mt5tk/internal/configurator/entities"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
)

func buildNew() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "create new MT5 server connection profile",
		Run: func(cmd *cobra.Command, args []string) {
			createNewProfile()
		},
	}

	return cmd
}

func createNewProfile() string {
	var profile model2.Profile

	var qs = []*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "Enter a new profile name:"},
			Validate: func(val interface{}) error {
				str, ok := val.(string)
				if !ok || str == "" {
					return errors.New("Name cannot be empty.")
				}
				if cfg.IfProfileExist(str) {
					return errors.New("A profile with such name does already exist.")
				}
				return nil
			},
		},
		{
			Name:   "serverAddress",
			Prompt: &survey.Input{Message: "Enter a server address (format: <IP address>:<port>):"},
			Validate: func(val interface{}) error {
				str, ok := val.(string)
				err := validate.Var(str, "tcp_addr")
				if !ok || err != nil {
					return errors.New("Server address must be a valid TCP address.")
				}
				return nil
			},
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
		{
			Name:     "skipVerifySsl",
			Prompt:   &survey.Confirm{Message: "Skip ssl verification:", Default: false},
			Validate: survey.Required,
		},
	}
	err := survey.Ask(qs, &profile)
	if err != nil {
		log.Fatal(err)
	}
	desc, err := yaml.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}
	confirmed := false
	prompt := &survey.Confirm{
		Message: "Following profile settings will be stored in the configuration file:\n" + string(desc) + "\nAre you sure?",
	}
	if err := survey.AskOne(prompt, &confirmed); err != nil {
		log.Fatal(err)
	}
	if confirmed {
		cfg.AddNewProfile(profile)
		log.Println("ok, profile successfully stored")
	} else {
		log.Println("ok, nothing has been changed")
	}

	return profile.Name
}
