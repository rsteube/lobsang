package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "manage projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectCmd).Standalone()


	rootCmd.AddCommand(projectCmd)
}
