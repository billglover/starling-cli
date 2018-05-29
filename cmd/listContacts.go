package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listContactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "List contacts",
	Run:   listContacts,
}

func init() {
	listCmd.AddCommand(listContactsCmd)
}

func listContacts(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	cons, _, err := sb.Contacts(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(*cons) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(*cons) {
		limit = len(*cons)
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%3s %30s %40s\n", "#", "Name", "UUID")
		for i := 0; i < limit; i++ {
			c := (*cons)[i]
			fmt.Printf("%s %30s %40s\n", color.BlueString("%03d", i), c.Name, c.UID)
		}
	} else {
		color.Green("%3s %30s\n", "#", "Name")
		for i := 0; i < limit; i++ {
			c := (*cons)[i]
			fmt.Printf("%s %30s\n", color.BlueString("%03d", i), c.Name)
		}
	}
}
