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

var createGoalCmd = &cobra.Command{
	Use:     "goal",
	Short:   "Create a savings goal",
	Aliases: []string{"g"},
	Run:     createGoal,
}

func init() {
	var name string
	createGoalCmd.Flags().StringVar(&name, "name", "", "name of the saving goal you want to create")
	viper.BindPFlag("name", createGoalCmd.Flags().Lookup("name"))
	createGoalCmd.MarkFlagRequired("name")

	var target float64
	createGoalCmd.Flags().Float64Var(&target, "target", 0.0, "target amount for the savings goal")
	viper.BindPFlag("target", createGoalCmd.Flags().Lookup("target"))

	createCmd.AddCommand(createGoalCmd)
}

func createGoal(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)

	uid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("unable to generate UID:", err)
		os.Exit(1)
	}

	amount := new(starling.Amount)
	if viper.GetFloat64("target") != 0.0 {
		amount.Currency = viper.GetString("currency")
		amount.MinorUnits = int64(viper.GetFloat64("target") * 100)
	}

	sgr := starling.SavingsGoalRequest{
		Name:     viper.GetString("name"),
		Currency: viper.GetString("currency"),
		Target:   *amount,
	}

	_, err = sb.CreateSavingsGoal(ctx, uid.String(), sgr)
	if err != nil {
		fmt.Println("unable to create savings goal:", err)
		os.Exit(1)
	}

	if viper.GetBool("uuid") == true {
		fmt.Println(uid)
	}
}
