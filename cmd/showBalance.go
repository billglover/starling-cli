package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var showBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Show a summary of your balance",
	Run:   showBalance,
	Args:  cobra.NoArgs,
}

func init() {
	showCmd.AddCommand(showBalanceCmd)
}

func showBalance(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	bal, _, err := sb.AccountBalance(ctx)
	check(err, "unable to show balance")

	key := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%-20s %10.2f\n", key("Amount:"), bal.Amount)
	fmt.Printf("%-20s %10.2f\n", key("Available:"), bal.Available)
	fmt.Printf("%-20s %10.2f\n", key("Cleared:"), bal.Cleared)
	fmt.Printf("%-20s %10.2f\n", key("Overdraft:"), bal.Overdraft)
	fmt.Printf("%-20s %10.2f\n", key("Pending:"), bal.PendingTxns)
	fmt.Printf("%-20s %10.2f\n", key("Effective:"), bal.Effective)
	fmt.Printf("%-20s %10s\n", key("Currency:"), bal.Currency)
}
