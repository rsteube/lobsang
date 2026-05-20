package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/lobsang/cmd/lobsang/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalCompletion(
		action.ActionProjects(),
		action.ActionDate(),
	)
}
