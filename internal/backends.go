package internal

import (
	"github.com/rsteube/lobsang/internal/backend"
	"github.com/rsteube/lobsang/internal/backend/toggl"
)

var Backends = make(map[string]backend.Backend)

func init() {
    Backends["toggl"] = &toggl.Toggl{}
}
