package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of items based on sub-command",
	Run:   list,
}

func init() {
	var limit int
	listCmd.PersistentFlags().IntVar(&limit, "limit", 10, "number of transactions to show")
	viper.BindPFlag("limit", listCmd.PersistentFlags().Lookup("limit"))

	rootCmd.AddCommand(listCmd)

	var from string
	var to string
	listCmd.PersistentFlags().StringVar(&from, "from", "", "filter results from this date (dd/mm/yyyy)")
	listCmd.PersistentFlags().StringVar(&to, "to", "", "filter results to this date (dd/mm/yyyy)")
}

func list(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("missing object, you need to list something e.g. list txns")
		os.Exit(1)
	}
}
