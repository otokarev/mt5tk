package entities

type Profile struct {
	Name          string `yaml:"name"`
	ServerAddress string `yaml:"serverAddress"`
	Login         string `yaml:"login"`
	Password      string `yaml:"password"`
}
