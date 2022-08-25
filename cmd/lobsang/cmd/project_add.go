package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var project_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_addCmd).Standalone()


	projectCmd.AddCommand(project_addCmd)
}
