package cmd

import (
	"github.com/otokarev/mt5tk/interanal/cmd/group"
	"github.com/otokarev/mt5tk/interanal/cmd/symbol"
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string
var server string
var login string
var password string
var skipVerifySsl bool

var modelFactory model.Factory

var RootCmd = &cobra.Command{
	Use:   "mt5tk",
	Short: "Utility to query MT5 server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initModelFactory)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mt5tk.yml)")
	RootCmd.PersistentFlags().StringVar(&server, "server", "", "MT5 server DSN")
	RootCmd.PersistentFlags().StringVar(&login, "login", "", "MT5 server login")
	RootCmd.PersistentFlags().StringVar(&password, "password", "", "MT5 server password")
	RootCmd.PersistentFlags().BoolVar(&skipVerifySsl, "skip_verify_ssl", false, "Skip SSL verification while communicate with MT5 server")

	viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("login", RootCmd.PersistentFlags().Lookup("login"))
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("skip_verify_ssl", RootCmd.PersistentFlags().Lookup("skip_verify_ssl"))

	RootCmd.AddCommand(symbol.Build(&modelFactory))
	RootCmd.AddCommand(group.Build(&modelFactory))
}

func initModelFactory() {
	conn := connection.NewConnection(
		viper.GetString("server"),
		viper.GetString("login"),
		viper.GetString("password"),
		viper.GetBool("skip_verify_ssl"))
	modelFactory = model.Factory{Client: &client.Client{Connection: conn}}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mt5tk.yml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
