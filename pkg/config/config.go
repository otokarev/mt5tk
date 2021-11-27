package config

type Config struct {
	Url string `yaml:"url" env:"SERVER" env-required`
	Login string `yaml:"login" env:"LOGIN" env-required`
	Password string `yaml:"password" env:"PASSWORD" env-required`
	SkipVerifySsl bool `yaml:"skip_verify_ssl" env:"SKIP_VERIFY_SSL" env-default:false`
}
