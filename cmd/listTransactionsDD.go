package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listTransactionsDDCmd = &cobra.Command{
	Use:   "dd",
	Short: "List direct-debit transactions",
	Run:   listTransactionsDD,
	Args:  cobra.NoArgs,
}

func init() {
	listTransactionsCmd.AddCommand(listTransactionsDDCmd)
}

func listTransactionsDD(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)

	txns, _, err := sb.DDTransactions(ctx, nil)
	check(err, "unable to list DD transactions")

	if len(txns) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(txns) {
		limit = len(txns)
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%3s %30s %10s %30s %40s\n", "#", "Created", "Amount", "Narrative", "UUID")
		for i := 0; i < limit; i++ {
			txn := txns[i]
			fmt.Printf("%s %30s %10.2f %30s %40s\n", color.BlueString("%03d", i), txn.Created, txn.Amount, txn.Narrative, txn.UID)
		}
	} else {
		color.Green("%3s %30s %10s %30s\n", "#", "Created", "Amount", "Narrative")
		for i := 0; i < limit; i++ {
			txn := txns[i]
			fmt.Printf("%s %30s %10.2f %30s\n", color.BlueString("%03d", i), txn.Created, txn.Amount, txn.Narrative)
		}
	}

	if limit < len(txns) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d transactions\n", limit, len(txns))
	}
}
