package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deleteContactCmd = &cobra.Command{
	Use:     "contact uid",
	Short:   "Delete a contact account",
	Aliases: []string{"c"},
	Run:     deleteContact,
	Args:    cobra.ExactArgs(1),
}

func init() {
	deleteCmd.AddCommand(deleteContactCmd)
}

func deleteContact(cmd *cobra.Command, args []string) {
	contactUID, err := uuid.Parse(args[0])
	if err != nil {
		fmt.Println("You must provide the UID for the contact you want to delete.")
		os.Exit(1)
	}

	ctx := context.Background()
	sb := newClient(ctx)

	resp, err := sb.DeleteContact(ctx, contactUID.String())
	if err != nil {
		fmt.Println("unable to delete contact:", resp.Status, err)
		os.Exit(1)
	}
}
