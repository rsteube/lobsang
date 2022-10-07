package backend

import "time"

type Backend interface {
	GetClients() (map[string]Client, error)
	GetProjects(client Client) (map[string]Project, error)
	GetTasks(project Project, from, to time.Time) ([]Task, error)
	Login() error
}

type Client struct {
	ID          string
	Name        string
	Description string
}

type Project struct {
	ID          string
	Name        string
	Description string
}

type Task struct {
	ID          string
	Description string
	Tags        []Tag
	From        time.Time
	To          time.Time
}

type Tag struct {
	ID          string
	Name        string
	Description string
	Color       string
}

type Auth interface{}
