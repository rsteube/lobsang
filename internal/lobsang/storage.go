package lobsang

import (
	"fmt"
	"os"
	"time"

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

func pathRange(from, to string) ([]string, error) {
	f, err := time.Parse("2006-01-02", from)
	if err != nil {
		return nil, err
	}

	if to == "" {
		to = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}

	t, err := time.Parse("2006-01-02", to)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0)
	for d := f; d.Before(t); d = d.AddDate(0, 0, 1) {
		paths = append(paths, d.Format("2006/01/2006-01-02.yaml"))
	}
	return paths, nil
}
