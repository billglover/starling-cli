package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display a table of information based on sub-command",
	Long: `Show is a command that queries the Starling Bank API and displays a 
table of information based on the associated sub-command. For example:

	starling-cli show account`,
	Run:  show,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(cmd *cobra.Command, args []string) {
	fmt.Printf("Error: invalid command \"%s\" provided\n", args[0])
	cmd.Usage()
	os.Exit(1)
}
