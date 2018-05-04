package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "Display a list of items",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"transactions", "contacts", "goals", "account"},
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
	case "account":
		listAccount()
	case "transactions":
		listTransactions()
	case "contacts":
		listContacts()
	case "goals":
		listGoals()
	}
}

func listTransactions() {
	ctx := context.Background()
	sb := newClient(ctx)

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
	ctx := context.Background()
	sb := newClient(ctx)
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
	ctx := context.Background()
	sb := newClient(ctx)
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

func listAccount() {
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
