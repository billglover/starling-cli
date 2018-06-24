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

var transferToCmd = &cobra.Command{
	Use:   "to",
	Short: "Transfer money to a savings goal",
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

	result, _, err := sb.AddMoney(ctx, goalID.String(), amt)
	check(err, "unable to transfer funds to savings goal")

	fmt.Println("transfer complete:", result)
}
