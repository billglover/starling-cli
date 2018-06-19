package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create things like goals, payees, etc.",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(createCmd)
}
