package action

import (
	"github.com/rsteube/carapace"
)

func ActionTags(project string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
	//	projects, err := storage.Projects()
	//	if err != nil {
	//		return carapace.ActionMessage(err.Error())
	//	}

	//	vals := make([]string, 0)
	//	for _, tag := range projects[project].Tags { // TODO check exists
	//		vals = append(vals, tag)
	//	}
	//	return carapace.ActionValues(vals...)
    return carapace.ActionValues()
	})
}
