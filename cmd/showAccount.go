package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var showAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Show an account summary",
	Run:   showAccount,
	Args:  cobra.NoArgs,
}

func init() {
	showCmd.AddCommand(showAccountCmd)
}

func showAccount(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	act, _, err := sb.Account(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%-20s %40s\n", key("Name:"), act.Name)
	fmt.Printf("%-20s %40s\n", key("Number:"), act.AccountNumber)
	fmt.Printf("%-20s %40s\n", key("Sort Code:"), act.SortCode)
	fmt.Printf("%-20s %40s\n", key("BIC:"), act.BIC)
	fmt.Printf("%-20s %40s\n", key("IBAN:"), act.IBAN)
	fmt.Printf("%-20s %40s\n", key("Currency:"), act.Currency)
	fmt.Printf("%-20s %40s\n", key("Created:"), act.CreatedAt)
	fmt.Printf("%-20s %40s\n", key("UID:"), act.UID)
}
