package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listFeedCmd = &cobra.Command{
	Use:   "feed",
	Short: "List recent feed items",
	Run:   listFeed,
}

func init() {
	listCmd.AddCommand(listFeedCmd)
}

func listFeed(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Printf("Error: invalid command \"%s\" provided\n", args[0])
		cmd.Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	sb := newClient(ctx)

	act := viper.GetString("account")
	cat := viper.GetString("category")

	items, _, err := sb.Feed(ctx, act, cat, nil)
	check(err, "unable to list feed")

	if len(items) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(items) {
		limit = len(items)
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%-3s %-35s %-10s %-42s %-40s\n", "#", "Time", "Amount", "Reference", "UUID")
		for i := 0; i < limit; i++ {
			item := items[i]
			fmt.Printf("%s %-35s %-10.2f %-42s %-40s\n", color.BlueString("%03d", i), item.TransactionTime, float64(item.Amount.MinorUnits)/100, item.Reference, item.FeedItemUID)
		}
	} else {
		color.Green("%-3s %-35s %-10s %-20s %-20s %-42s\n", "#", "Created", "Amount", "Source", "Source Sub-Type", "Reference")
		for i := 0; i < limit; i++ {
			item := items[i]
			fmt.Printf("%s %-35s %-10.2f %-20s %-20s %-42s\n", color.BlueString("%03d", i), item.TransactionTime, float64(item.Amount.MinorUnits)/100, item.Source, item.SourceSubType, item.Reference)
		}
	}

	if limit < len(items) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d transactions\n", limit, len(items))
	}
}
