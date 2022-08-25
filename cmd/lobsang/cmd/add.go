package cmd

import (
	"github.com/rsteube/carapace"
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
		action.ActionDate(),
		action.ActionDuration(),
	)
}
