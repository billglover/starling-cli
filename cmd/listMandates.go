package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listMandatesCmd = &cobra.Command{
	Use:   "mandates",
	Short: "List Direct Debit mandates",
	Run:   listMandates,
	Args:  cobra.NoArgs,
}

func init() {
	listCmd.AddCommand(listMandatesCmd)
}

func listMandates(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	ms, _, err := sb.DirectDebitMandates(ctx)
	check(err, "unable to list mandates")

	if len(ms) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(ms) {
		limit = len(ms)
	}

	color.Green("%03s %-30s %-30s %-20s %-10s\n", "  #", "Reference", "Created", "Orignator", "Status")
	for i := 0; i < limit; i++ {
		m := ms[i]
		fmt.Printf("%s %-30s %-30s %-20s %-10s\n", color.BlueString("%03d", i), m.Reference, m.Created, m.OriginatorName, m.Status)
	}

	if limit < len(ms) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d mandates\n", limit, len(ms))
	}
}
