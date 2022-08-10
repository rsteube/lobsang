package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/lobsang/internal/storage"
)

func ActionProjects() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		projects, err := storage.Projects()
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
