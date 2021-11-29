package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jessevdk/go-flags"
	"github.com/otokarev/mt5tk/interanal/cli/group"
	"github.com/otokarev/mt5tk/interanal/cli/symbol"
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/config"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model"
	"log"
	"os"
)

var cfg config.Config

type CliOptions struct {
}

var cliOptions CliOptions

var cliParser = flags.NewParser(&cliOptions, flags.Default)

func main() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	conn := connection.NewConnection(cfg.Url, cfg.Login, cfg.Password, cfg.SkipVerifySsl)
	factory := &model.Factory{Client: &client.Client{Connection: conn}}

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("Bad command")
	}
	ns := args[0]
	args = args[1:]

	switch ns {
	case "symbol":
		symbol.Parse(factory, cliParser, args)
	case "group":
		group.Parse(factory, cliParser, args)
	default:
		log.Fatal("Bad command")
	}
}
