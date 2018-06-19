package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listGoalsCmd = &cobra.Command{
	Use:   "goals",
	Short: "List savings goals",
	Run:   listGoals,
	Args:  cobra.NoArgs,
}

func init() {
	listCmd.AddCommand(listGoalsCmd)
}

func listGoals(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	sb := newClient(ctx)
	goals, _, err := sb.SavingsGoals(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(goals) == 0 {
		return
	}

	limit := viper.GetInt("limit")
	if limit > len(goals) {
		limit = len(goals)
	}

	uuid := viper.GetBool("uuid")

	if uuid == true {
		color.Green("%-3s %-20s %-10s %-10s %-11s %-40s\n", "#", "Name", "Saved", "Target", "Percent", "UID")
		for i := 0; i < limit; i++ {
			g := goals[i]
			saved := float64(g.TotalSaved.MinorUnits) / 100
			target := float64(g.Target.MinorUnits) / 100
			fmt.Printf("%-s %-20s %10.2f %10.2f %10d%% %-40s\n", color.BlueString("%03d", i), g.Name, saved, target, g.SavedPercentage, g.UID)
		}
	} else {
		color.Green("%3s %-20s %10s %10s %11s\n", "#", "Name", "Saved", "Target", "Percentage")
		for i := 0; i < limit; i++ {
			g := goals[i]
			saved := float64(g.TotalSaved.MinorUnits) / 100
			target := float64(g.Target.MinorUnits) / 100
			fmt.Printf("%s %-20s %10.2f %10.2f %10d%%\n", color.BlueString("%03d", i), g.Name, saved, target, g.SavedPercentage)
		}
	}

	if limit < len(goals) {
		color.Set(color.FgHiMagenta)
		fmt.Printf("%d of %d savings goals\n", limit, len(goals))
	}
}
