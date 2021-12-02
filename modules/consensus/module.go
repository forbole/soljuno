package consensus

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module implements the consensus utils
type Module struct {
	db db.ConsensusDb
}

// NewModule builds a new Module instance
func NewModule(db db.ConsensusDb) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "consensus"
}
