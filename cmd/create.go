package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create things like goals, payees, etc.",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(1),
	Run:     create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(cmd *cobra.Command, args []string) {
	fmt.Printf("Error: invalid command \"%s\" provided\n\n", args[0])
	cmd.Usage()
	os.Exit(1)
}
