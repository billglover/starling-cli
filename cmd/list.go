package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	"github.com/billglover/starling"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "A brief description of your command",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"transactions", "contacts"},
	Run: func(cmd *cobra.Command, args []string) {
		listTransactions()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTransactions() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: viper.GetString("token")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	client := starling.NewClient(tc)

	txns, _, _ := client.Transactions(ctx, nil)

	for i, txn := range *txns {
		fmt.Println(i, txn.Created, txn.Amount, txn.Currency, txn.Narrative)
	}
}
