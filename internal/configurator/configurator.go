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
	cfg   entities.Configurations
	viper *viper.Viper
}

func LoadConfig(fileName string, fileType string, path string) *Configuration {
	c := new(Configuration)
	c.viper = viper.New()
	c.viper.AddConfigPath(path)
	c.viper.SetConfigType(fileType)
	c.viper.SetConfigName(fileName)

	c.readConfig()

	return c
}

func LoadDefaultConfig() *Configuration {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	c := new(Configuration)
	c.viper = viper.New()
	c.viper.AddConfigPath(home)
	c.viper.SetConfigType("yaml")
	c.viper.SetConfigName(".mt5tk.yml")
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

func (c *Configuration) IfProfileExist(name string) bool {
	p := c.GetProfileConfig(name)

	if p.Name == "" {
		return false
	} else {
		return true
	}
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

func (c *Configuration) SwitchToProfile(name string) {
	c.cfg.Profile = name
	c.storeConfig()
}

func (c *Configuration) AddNewProfile(profile entities.Profile) {
	c.cfg.Profiles = append(c.cfg.Profiles, profile)
	c.storeConfig()
}

func (c *Configuration) DeleteProfile(name string) {
	for i, p := range c.cfg.Profiles {
		if name == p.Name {
			c.cfg.Profiles = append(c.cfg.Profiles[:i], c.cfg.Profiles[i+1:]...)
		}
	}
	c.storeConfig()
}

func (c *Configuration) GetProfileConfig(name string) entities.Profile {
	for _, v := range c.cfg.Profiles {
		if v.Name == name {
			return v
		}
	}

	return entities.Profile{}
}

func (c *Configuration) GetProfileConfigYaml(name string) string {
	p := c.GetProfileConfig(name)
	if p.Name == "" {
		log.Fatalf("profile %s does not exist", name)
	}

	b, err := yaml.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func (c *Configuration) readConfig() {

	err := c.viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config file, error: %v \n", err)
	}

	err = c.viper.Unmarshal(&c.cfg)
	if err != nil {
		log.Fatalf("cannot parse config file, error: %v \n", err)
	}
}

func (c *Configuration) storeConfig() {
	data, err := yaml.Marshal(c.cfg)
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
	err = c.viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
	err = c.viper.WriteConfig()
	if err != nil {
		log.Fatalf("cannot store config file, error: %v \n", err)
	}
}
