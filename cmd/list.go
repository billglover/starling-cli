package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/billglover/starling"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "Display a list of items",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"transactions", "contacts", "goals"},
	Run:       list,
}

func init() {
	rootCmd.AddCommand(listCmd)

	var from string
	var to string

	listCmd.PersistentFlags().StringVar(&from, "from", "", "filter results from this date (dd/mm/yyyy)")
	listCmd.PersistentFlags().StringVar(&to, "to", "", "filter results to this date (dd/mm/yyyy)")
}

func list(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("missing object, you need to list something e.g. list transactions")
		os.Exit(1)
	}

	switch args[0] {
	case "transactions":
		listTransactions()
	case "contacts":
		listContacts()
	case "goals":
		listGoals()
	}
}

func listTransactions() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: viper.GetString("token")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	sb := starling.NewClient(tc)
	txns, _, err := sb.Transactions(ctx, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, txn := range *txns {
		fmt.Printf("%s %s %10.2f %s\n", color.BlueString("%03d", i), txn.Created, txn.Amount, txn.Narrative)
	}
}

func listContacts() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: viper.GetString("token")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	sb := starling.NewClient(tc)
	cons, _, err := sb.Contacts(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, c := range *cons {
		fmt.Printf("%s %s\n", color.BlueString("%03d", i), c.Name)
	}
}

func listGoals() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: viper.GetString("token")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	sb := starling.NewClient(tc)
	goals, _, err := sb.SavingsGoals(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, g := range *goals {
		saved := float64(g.TotalSaved.MinorUnits) / 100
		target := float64(g.Target.MinorUnits) / 100
		fmt.Printf("%s %-20s %10.2f %10.2f %10d\n", color.BlueString("%03d", i), g.Name, saved, target, g.SavedPercentage)
	}
}
