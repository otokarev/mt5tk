package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/otokarev/mt5tk/pkg/config"
	"github.com/otokarev/mt5tk/pkg/connection"
	"log"
)

var cfg config.Config

func main() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	conn := connection.NewConnection(cfg.Url, cfg.Login, cfg.Password, cfg.SkipVerifySsl)

	conn.Ping()
}
