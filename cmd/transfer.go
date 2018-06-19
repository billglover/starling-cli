package cmd

import (
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer money to/from a savings goal",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(transferCmd)
}
