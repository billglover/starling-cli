package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var token string

var rootCmd = &cobra.Command{
	Use:   "starling-cli",
	Short: "A command line interface to Starling Bank",
	Long: `This is a basic command line interface for personal Starling 
Bank accounts. It allows you to perform basic banking from 
the command line. For example:

	starling-cli list transactions

The Starling API is still under active development and until it
stabilises there may be some instability in the functionality 
provided`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starling.yaml)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "API access token")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".starling")
	}

	viper.SetEnvPrefix("starling")
	viper.AutomaticEnv()

	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
