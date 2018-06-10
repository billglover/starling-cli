package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/billglover/starling"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createContactCmd = &cobra.Command{
	Use:     "contact",
	Short:   "Create a contact account",
	Aliases: []string{"c"},
	Run:     createContact,
}

func init() {
	var name string
	createContactCmd.Flags().StringVar(&name, "name", "", "name of the contact account that you want to create")
	viper.BindPFlag("cname", createContactCmd.Flags().Lookup("name"))
	createContactCmd.MarkFlagRequired("cname")

	var an string
	createContactCmd.Flags().StringVar(&an, "account-number", "", "the account number of the contact account")
	viper.BindPFlag("account-number", createContactCmd.Flags().Lookup("account-number"))
	createContactCmd.MarkFlagRequired("account-number")

	var sc string
	createContactCmd.Flags().StringVar(&sc, "sort-code", "", "the sort code of the contact account")
	viper.BindPFlag("sort-code", createContactCmd.Flags().Lookup("sort-code"))
	createContactCmd.MarkFlagRequired("sort-code")

	createCmd.AddCommand(createContactCmd)
}

func createContact(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)

	uid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("unable to generate UID:", err)
		os.Exit(1)
	}

	ca := starling.ContactAccount{
		UID:           uid.String(),
		Type:          "UK_ACCOUNT_AND_SORT_CODE",
		Name:          viper.GetString("cname"),
		AccountNumber: viper.GetString("account-number"),
		SortCode:      viper.GetString("sort-code"),
	}

	_, err = sb.CreateContactAccount(ctx, ca)
	if err != nil {
		fmt.Println("unable to create contact account:", err)
		os.Exit(1)
	}

	if viper.GetBool("uuid") == true {
		fmt.Println(uid)
	}
}
