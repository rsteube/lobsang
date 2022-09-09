package lobsang

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type project struct {
	Description string
	Tags        []string
}

func Location() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("could not determine userhome")
	}
	// TODO different path for other os
	return fmt.Sprintf("%v/.local/share/lobsang/", home)
}

func Projects() (map[string]project, error) {
	content, err := os.ReadFile(Location() + "projects.yaml")
	if err != nil {
		return nil, err
	}

	var projects map[string]project
	if err := yaml.Unmarshal(content, &projects); err != nil {
		return nil, err
	}
	return projects, nil

}
