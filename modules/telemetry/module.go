package telemetry

import (
	"github.com/forbole/soljuno/modules"
)

const (
	ModuleName = "telemetry"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

// Module represents the telemetry module
type Module struct {
	cfg Config
}

// NewModule returns a new Module implementation
func NewModule(cfg Config) *Module {
	return &Module{
		cfg: cfg,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return ModuleName
}

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	return RunAdditionalOperations(m.cfg)
}
