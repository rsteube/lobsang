package toggl

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rsteube/lobsang/internal/backend"
)

type Toggl struct {
	Token string
}

// GetClients implements backend.Backend
func (*Toggl) GetClients() (map[string]backend.Client, error) {
	panic("unimplemented")
}

// GetProjects implements backend.Backend
func (*Toggl) GetProjects(client backend.Client) (map[string]backend.Project, error) {
	panic("unimplemented")
}

// GetTasks implements backend.Backend
func (*Toggl) GetTasks(project backend.Project, from time.Time, to time.Time) ([]backend.Task, error) {
	panic("unimplemented")
}

// Login implements backend.Backend
func (t *Toggl) Login() error {
	i := textinput.New()

	p := tea.NewProgram(i, tea.WithOutput(os.Stderr))
	tm, err := p.StartReturningModel()
	if err != nil {
		return fmt.Errorf("failed to run input: %w", err)
	}
	m := tm.(model)

	fmt.Println(m.textinput.Value())
	panic("unimplemented")
}
