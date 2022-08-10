package cmd

import (
	"unicode"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/lobsang/cmd/lobsang/cmd/action"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	addCmd.Flags().String("tags", "", "tags for entry")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"tags": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionMessage("missing project")
			}
			return action.ActionTags(c.Args[0]).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		action.ActionProjects(),
		ActionDate(),
		ActionDuration(),
	)
}

func ActionDate() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(
			carapace.ActionValues("monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday", "today", "yesterday").Style(style.Blue),
		)

		if len(c.CallbackValue) > 0 && unicode.IsDigit([]rune(c.CallbackValue)[0]) {
			batch = append(batch, time.ActionDate())
		}

		return batch.ToA()
	})
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
