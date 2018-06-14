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
	Run:     delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("missing object, you need to delete something e.g. goal")
		os.Exit(1)
	}
}
