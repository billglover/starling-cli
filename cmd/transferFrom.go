package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var transferFromCmd = &cobra.Command{
	Use:   "from",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transferFrom called")
	},
}

func init() {
	transferCmd.AddCommand(transferFromCmd)
}
