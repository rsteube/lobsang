package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/lobsang/internal"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Backends["toggl"].Login()
	},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	authCmd.AddCommand(auth_loginCmd)
}
