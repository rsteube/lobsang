package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalCompletion(
		carapace.ActionValues("TODO-projectA", "TODO-projectB"),
		ActionDate(),
		ActionDuration(),
	)
}

func ActionDate() carapace.Action {
	return carapace.Batch(
		carapace.ActionValues("monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"),
		time.ActionDate(),
	).ToA()
}

func ActionDuration() carapace.Action {
	return carapace.ActionValues(
		"0.25", "0.5", "0.75", "1.0", "1.25", "1.5", "1.75", "2.0",
		"2.25", "2.5", "2.75", "3.0", "3.25", "3.5", "3.75", "4.0",
		"4.25", "4.5", "4.75", "5.0", "5.25", "5.5", "5.75", "6.0",
		"6.25", "6.5", "6.75", "7.0", "7.25", "7.5", "7.75", "8.0",
		"9.25", "9.5", "9.75",
	)
}
