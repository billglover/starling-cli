package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "Display a list of items based on sub-command",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"transactions", "contacts", "goals", "mandates", "addresses", "payments"},
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
	case "contacts":
		listContacts()
	case "goals":
		listGoals()
	case "mandates":
		listMandates()
	case "addresses":
		listAddresses()
	case "payments":
		listPayments()
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

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%3s %30s %40s\n", "#", "Name", "UUID")
		for i, c := range *cons {
			fmt.Printf("%s %30s %40s\n", color.BlueString("%03d", i), c.Name, c.UID)
		}
	} else {
		color.Green("%3s %30s\n", "#", "Name")
		for i, c := range *cons {
			fmt.Printf("%s %30s\n", color.BlueString("%03d", i), c.Name)
		}
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

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%s %-20s %10s %10s %11s %40s\n", "  #", "Name", "Saved", "Target", "Percentage", "UID")
		for i, g := range *goals {
			saved := float64(g.TotalSaved.MinorUnits) / 100
			target := float64(g.Target.MinorUnits) / 100
			fmt.Printf("%s %-20s %10.2f %10.2f %10d%% %40s\n", color.BlueString("%03d", i), g.Name, saved, target, g.SavedPercentage, g.UID)
		}
	} else {
		color.Green("%s %-20s %10s %10s %11s\n", "  #", "Name", "Saved", "Target", "Percentage")
		for i, g := range *goals {
			saved := float64(g.TotalSaved.MinorUnits) / 100
			target := float64(g.Target.MinorUnits) / 100
			fmt.Printf("%s %-20s %10.2f %10.2f %10d%%\n", color.BlueString("%03d", i), g.Name, saved, target, g.SavedPercentage)
		}
	}
}

func listMandates() {
	ctx := context.Background()
	sb := newClient(ctx)
	ms, _, err := sb.DirectDebitMandates(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	color.Green("%03s %-30s %-30s %-20s %-10s\n", "  #", "Reference", "Created", "Orignator", "Status")
	for i, m := range ms {
		fmt.Printf("%s %-30s %-30s %-20s %-10s\n", color.BlueString("%03d", i), m.Reference, m.Created, m.OriginatorName, m.Status)
	}
}

func listAddresses() {
	ctx := context.Background()
	sb := newClient(ctx)
	addrs, _, err := sb.AddressHistory(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cur := addrs.Current

	key := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%-20s %20s\n", key("Street:"), cur.Street)
	fmt.Printf("%-20s %20s\n", key("City:"), cur.City)
	fmt.Printf("%-20s %20s\n", key("Country:"), cur.Country)
	fmt.Printf("%-20s %20s\n", key("Postcode:"), cur.Postcode)

	if len(addrs.Previous) == 0 {
		return
	}

	color.Green("%03s %-30s %-20s %-20s %-10s\n", "  #", "Street", "City", "Country", "Postcode")
	for i, a := range addrs.Previous {
		fmt.Printf("%s %-30s %-20s %-20s %-10s\n", color.BlueString("%03d", i), a.Street, a.City, a.Country, a.Postcode)
	}
}

func listPayments() {
	ctx := context.Background()
	sb := newClient(ctx)
	pos, _, err := sb.ScheduledPayments(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(pos) == 0 {
		return
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%-03s %-30s %-20s %-20s %-10s %-10s %-20s %-40s\n", "#", "Recipient", "Reference", "Next Payment", "Amount", "Currency", "Recurrence", "UUID")
		for i, po := range pos {
			fmt.Printf("%-s %-30s %-20s %-20s %-10.2f %-10s %-20s %-40s\n", color.BlueString("%03d", i), po.RecipientName, po.Reference, po.NextDate, po.Amount, po.Currency, po.RecurrenceRule.Frequency, po.UID)
		}
	} else {
		color.Green("%-03s %-30s %-20s %-20s %-10s %-10s %-20s\n", "#", "Recipient", "Reference", "Next Payment", "Amount", "Currency", "Recurrence")
		for i, po := range pos {
			fmt.Printf("%s %-30s %-20s %-20s %-10.2f %-10s %-20s\n", color.BlueString("%03d", i), po.RecipientName, po.Reference, po.NextDate, po.Amount, po.Currency, po.RecurrenceRule.Frequency)
		}
	}
}
