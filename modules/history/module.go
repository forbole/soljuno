package history

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
