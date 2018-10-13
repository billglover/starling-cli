package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listAddressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "List addresses",
	Aliases: []string{"a"},
	Run:   listAddresses,
	Args:  cobra.NoArgs,
}

func init() {
	listCmd.AddCommand(listAddressesCmd)
}

func listAddresses(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	addrs, _, err := sb.AddressHistory(ctx)
	check(err, "unable to list addresses")

	cur := addrs.Current

	key := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%-20s %20s\n", key("Street:"), cur.Street)
	fmt.Printf("%-20s %20s\n", key("City:"), cur.City)
	fmt.Printf("%-20s %20s\n", key("Country:"), cur.Country)
	fmt.Printf("%-20s %20s\n", key("Postcode:"), cur.Postcode)

	if len(addrs.Previous) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(addrs.Previous) {
		limit = len(addrs.Previous)
	}

	color.Green("%03s %-30s %-20s %-20s %-10s\n", "  #", "Street", "City", "Country", "Postcode")
	for i := 0; i < limit; i++ {
		a := addrs.Previous[i]
		fmt.Printf("%s %-30s %-20s %-20s %-10s\n", color.BlueString("%03d", i), a.Street, a.City, a.Country, a.Postcode)
	}

	if limit < len(addrs.Previous) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d previous addresses\n", limit, len(addrs.Previous))
	}
}
