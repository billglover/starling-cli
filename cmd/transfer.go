package cmd

import (
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer money to/from a savings goal",
}

func init() {
	rootCmd.AddCommand(transferCmd)
}
