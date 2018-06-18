package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/billglover/starling"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transferFromCmd = &cobra.Command{
	Use:   "from",
	Short: "Transfer money from a savings goal",
	Args:  cobra.ExactArgs(2),
	Run:   transferFrom,
}

func init() {
	transferCmd.AddCommand(transferFromCmd)
}

func transferFrom(cmd *cobra.Command, args []string) {

	goalID, err := uuid.Parse(args[0])
	if err != nil {
		fmt.Println("You must provide the UID for the savings goal.")
		os.Exit(1)
	}

	val, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("You must specify the amount you wish to transfer.")
		os.Exit(1)
	}

	amt := starling.Amount{
		MinorUnits: int64(100 * val),
		Currency:   viper.GetString("currency"),
	}

	ctx := context.Background()
	sb := newClient(ctx)

	result, _, err := sb.Withdraw(ctx, goalID.String(), amt)
	if err != nil {
		fmt.Println("Unable to transfer funds from the savings goal:", err)
		os.Exit(1)
	}

	fmt.Println("Transfer complete:", result)
}
