package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete things like goals, payees, etc.",
	Aliases: []string{"d"},
	Args:    cobra.MinimumNArgs(1),
	Run:     delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) {
	fmt.Printf("Error: invalid command \"%s\" provided\n", args[0])
	cmd.Usage()
	os.Exit(1)
}
