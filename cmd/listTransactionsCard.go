package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listTransactionsCardCmd = &cobra.Command{
	Use:   "card",
	Short: "List card transactions",
	Run:   listTransactionsCard,
}

func init() {
	listTransactionsCmd.AddCommand(listTransactionsCardCmd)
}

func listTransactionsCard(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)

	txns, _, err := sb.MastercardTransactions(ctx, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(*txns) == 0 {
		return
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%-3s %-30s %-10s %-30s %-40s\n", "#", "Created", "Amount", "Narrative", "UUID")
		for i, txn := range *txns {
			fmt.Printf("%-s %-30s %-10.2f %-30s %-40s\n", color.BlueString("%03d", i), txn.Created, txn.Amount, txn.Narrative, txn.UID)
		}
	} else {
		color.Green("%-3s %-30s %-10s %-30s\n", "#", "Created", "Amount", "Narrative")
		for i, txn := range *txns {
			fmt.Printf("%-s %-30s %-10.2f %-30s\n", color.BlueString("%03d", i), txn.Created, txn.Amount, txn.Narrative)
		}
	}
}
