package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/lobsang/internal/backend"
)

func ActionProjects(b backend.Backend, client backend.Client) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		projects, err := b.GetProjects(client)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, project := range projects {
			vals = append(vals, name, project.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
