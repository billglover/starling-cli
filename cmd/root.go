package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/billglover/starling"
	"golang.org/x/oauth2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
	var token string
	var env string
	var uuid bool

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starling.yaml)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "API access token")
	rootCmd.PersistentFlags().StringVar(&env, "env", "sandbox", "the environment you want to use: live, sandbox (default is sandbox)")
	rootCmd.PersistentFlags().BoolVar(&uuid, "uuid", false, "display UUID for objects")
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
	viper.BindPFlag("env", rootCmd.Flags().Lookup("env"))
	viper.BindPFlag("uuid", rootCmd.Flags().Lookup("uuid"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to read config file:", viper.ConfigFileUsed())
	}

	switch viper.GetString("env") {
	case "live":
		viper.Set("url", starling.ProdURL)
	case "sandbox":
		viper.Set("url", starling.SandboxURL)
	default:
		fmt.Printf("unrecognised environment specified '%s', expected 'sandbox' or 'live'\n", viper.GetString("env"))
		os.Exit(1)
	}
}

func newClient(ctx context.Context) *starling.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: viper.GetString("token")})
	tc := oauth2.NewClient(ctx, ts)

	baseURL, _ := url.Parse(viper.GetString("url"))
	opts := starling.ClientOptions{BaseURL: baseURL}
	return starling.NewClientWithOptions(tc, opts)
}
