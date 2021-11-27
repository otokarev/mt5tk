package main

import (
	"encoding/json"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/config"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model"
	"log"
)

var cfg config.Config

func main() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	conn := connection.NewConnection(cfg.Url, cfg.Login, cfg.Password, cfg.SkipVerifySsl)
	client := &client.Client{Connection: conn}
	factory := &model.Factory{Client: client}

	fmt.Println(factory.Symbol().List())

	json, err := json.MarshalIndent(factory.Symbol().Get("SPOT"), "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
