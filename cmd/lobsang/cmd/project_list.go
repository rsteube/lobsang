package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/lobsang/internal/storage"
	"github.com/spf13/cobra"
)

var project_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := storage.Projects()
		if err != nil {
			return err
		}
		for name, project := range projects {
			fmt.Printf("%v: %v\n", name, project.Description)
		}
		return nil
	},
}

func init() {
	carapace.Gen(project_listCmd).Standalone()

	projectCmd.AddCommand(project_listCmd)
}
