package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of items based on sub-command",
	Aliases: []string{"l"},
}

func init() {
	var limit int
	listCmd.PersistentFlags().IntVar(&limit, "limit", 10, "number of transactions to show")
	viper.BindPFlag("limit", listCmd.PersistentFlags().Lookup("limit"))

	rootCmd.AddCommand(listCmd)
}
