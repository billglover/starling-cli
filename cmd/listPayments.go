package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listPaymentsCmd = &cobra.Command{
	Use:   "payments",
	Short: "List payments",
	Run:   listPayments,
}

func init() {
	listCmd.AddCommand(listPaymentsCmd)
}

func listPayments(cmd *cobra.Command, args []string) {
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

	limit := viper.GetInt("limit")
	if limit > len(pos) {
		limit = len(pos)
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%-03s %-30s %-20s %-20s %-10s %-10s %-20s %-40s\n", "#", "Recipient", "Reference", "Next Payment", "Amount", "Currency", "Recurrence", "UUID")
		for i := 0; i < limit; i++ {
			po := pos[i]
			fmt.Printf("%-s %-30s %-20s %-20s %-10.2f %-10s %-20s %-40s\n", color.BlueString("%03d", i), po.RecipientName, po.Reference, po.NextDate, po.Amount, po.Currency, po.RecurrenceRule.Frequency, po.UID)
		}
	} else {
		color.Green("%-03s %-30s %-20s %-20s %-10s %-10s %-20s\n", "#", "Recipient", "Reference", "Next Payment", "Amount", "Currency", "Recurrence")
		for i := 0; i < limit; i++ {
			po := pos[i]
			fmt.Printf("%s %-30s %-20s %-20s %-10.2f %-10s %-20s\n", color.BlueString("%03d", i), po.RecipientName, po.Reference, po.NextDate, po.Amount, po.Currency, po.RecurrenceRule.Frequency)
		}
	}

	if limit < len(pos) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d payments\n", limit, len(pos))
	}
}
