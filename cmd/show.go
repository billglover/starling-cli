package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display a table of information based on sub-command",
	Long: `Show is a command that queries the Starling Bank API and displays a 
table of information based on the associated sub-command. For example:

	starling-cli show account`,
}

func init() {
	rootCmd.AddCommand(showCmd)
}
