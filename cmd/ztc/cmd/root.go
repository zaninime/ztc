package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zaninime/ztc/api"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "ztc",
	Short: "CLI tool for managing a ZeroTier Controller",
	Long: `ztc is a CLI tool for managing a ZeroTier Controller.

This tool will help you call the right API endpoints through semantic
commands.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ztc.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ztc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ztc")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("ztc")
	viper.SetDefault("baseUrl", "http://localhost:9993")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getAPIController() api.Controller {
	url, err := url.Parse(viper.GetString("baseUrl"))

	if err != nil {
		log.Fatal(err)
	}

	cntrl := api.Controller{
		BaseURL:   *url,
		AuthToken: viper.GetString("authToken"),
	}

	return cntrl
}
