package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var project_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_removeCmd).Standalone()


	projectCmd.AddCommand(project_removeCmd)
}
