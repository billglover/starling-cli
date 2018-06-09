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
	Run:     create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("missing object, you need to create something e.g. goal")
		os.Exit(1)
	}
}
