package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete things like goals, payees, etc.",
	Aliases: []string{"d"},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
