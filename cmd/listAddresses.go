package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listAddressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "List addresses",
	Run:   listAddresses,
}

func init() {
	listCmd.AddCommand(listAddressesCmd)
}

func listAddresses(cmd *cobra.Command, args []string) {
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

	limit := viper.GetInt("limit")
	if limit > len(addrs.Previous) {
		limit = len(addrs.Previous)
	}

	color.Green("%03s %-30s %-20s %-20s %-10s\n", "  #", "Street", "City", "Country", "Postcode")
	for i := 0; i < limit; i++ {
		a := addrs.Previous[i]
		fmt.Printf("%s %-30s %-20s %-20s %-10s\n", color.BlueString("%03d", i), a.Street, a.City, a.Country, a.Postcode)
	}
}
