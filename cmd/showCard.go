package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var showCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Show your card details",
	Run:   showCard,
	Args:  cobra.NoArgs,
}

func init() {
	showCmd.AddCommand(showCardCmd)
}

func showCard(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	c, _, err := sb.Card(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%-20s %40s\n", key("Type:"), c.Type)
	fmt.Printf("%-20s %40s\n", key("Name:"), c.NameOnCard)
	fmt.Printf("%-20s %40s\n", key("Card Num:"), "**** **** **** "+c.LastFourDigits)
	fmt.Printf("%-20s %40v\n", key("Enabled:"), c.Enabled)
	fmt.Printf("%-20s %40v\n", key("Activated:"), c.Activated)
	fmt.Printf("%-20s %40v\n", key("Cancelled:"), c.Cancelled)
	fmt.Printf("%-20s %40s\n", key("UID:"), c.UID)
}
