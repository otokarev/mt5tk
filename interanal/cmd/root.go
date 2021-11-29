package cmd

import (
	"github.com/otokarev/mt5tk/pkg/client"
	"github.com/otokarev/mt5tk/pkg/connection"
	"github.com/otokarev/mt5tk/pkg/model"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var server string
var login string
var password string
var skipVerifySsl bool

var modelFactory *model.Factory

var rootCmd = &cobra.Command{
	Use:   "mt5tk",
	Short: "Utility to query MT5 server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initModelFactory)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mt5tk.yml)")
	rootCmd.PersistentFlags().StringVar(&server, "server", "", "MT5 server DSN")
	rootCmd.PersistentFlags().StringVar(&login, "login", "", "MT5 server login")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "MT5 server password")
	rootCmd.PersistentFlags().BoolVar(&skipVerifySsl, "skip_verify_ssl", false, "Skip SSL verification while communicate with MT5 server")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("login", rootCmd.PersistentFlags().Lookup("login"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("skip_verify_ssl", rootCmd.PersistentFlags().Lookup("skip_verify_ssl"))
}

func initModelFactory() {
	conn := connection.NewConnection(
		viper.GetString("server"),
		viper.GetString("login"),
		viper.GetString("password"),
		viper.GetBool("skip_verify_ssl"))
	modelFactory = &model.Factory{Client: &client.Client{Connection: conn}}
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

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}
