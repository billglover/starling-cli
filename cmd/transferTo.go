package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var transferToCmd = &cobra.Command{
	Use:   "to",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	Run:   transferTo,
}

func init() {
	transferCmd.AddCommand(transferToCmd)
}

func transferTo(cmd *cobra.Command, args []string) {

	goalID, err := uuid.Parse(args[0])
	if err != nil {
		fmt.Println("You must provide the UID for the savings goal.")
		os.Exit(1)
	}

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("You must specify the amount you wish to transfer.")
		os.Exit(1)
	}

	fmt.Println("goal:", goalID)
	fmt.Println("amount:", amount)
}
