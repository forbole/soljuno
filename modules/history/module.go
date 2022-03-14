package history

import "github.com/forbole/soljuno/modules"

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	services []HistroyService
}

func NewModule() *Module {
	return &Module{}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "history"
}
