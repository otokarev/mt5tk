package configurator

import (
	"bytes"
	"github.com/otokarev/mt5tk/internal/configurator/entities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Configuration struct {
	cfg entities.Configurations
}

func LoadConfig(fileName string, fileType string, path string) *Configuration {
	viper.AddConfigPath(path)
	viper.SetConfigType(fileType)
	viper.SetConfigName(fileName)

	c := new(Configuration)
	c.readConfig()

	return c
}

func LoadDefaultConfig() *Configuration {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".mt5tk.yml")

	c := new(Configuration)
	c.readConfig()

	return c
}

func (c *Configuration) GetProfileNames() []string {
	var ret []string
	for _, a := range c.cfg.Profiles {
		ret = append(ret, a.Name)
	}

	return ret
}

func (c *Configuration) GetDefaultProfile() string {
	profiles := c.GetProfileNames()
	current := c.cfg.Profile

	if len(profiles) == 0 {
		panic("no profiles configured")
	}
	for _, p := range profiles {
		if current == p {
			return current
		}
	}

	// bad default profile set, lets recover it
	c.cfg.Profile = profiles[0]
	c.storeConfig()

	return c.cfg.Profile
}

func (c *Configuration) SwitchToProfile(profile string) {
	c.cfg.Profile = profile
	c.storeConfig()
}

func (c *Configuration) readConfig() {

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config file, error: %v \n", err)
	}

	err = viper.Unmarshal(&c.cfg)
	if err != nil {
		log.Fatalf("cannot parse config file, error: %v \n", err)
	}
}

func (c *Configuration) storeConfig() {
	data, err := yaml.Marshal(c.cfg)
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
	err = viper.WriteConfig()
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
}
