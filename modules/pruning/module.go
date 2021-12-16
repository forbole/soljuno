package pruning

import (
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var _ modules.Module = &Module{}

// Module represents the pruning module allowing to clean the database periodically
type Module struct {
	cfg    types.PruningConfig
	db     db.Database
	logger logging.Logger
	signal chan bool
}

// NewModule builds a new Module instance
func NewModule(cfg types.PruningConfig, db db.Database, logger logging.Logger) *Module {
	return &Module{
		cfg:    cfg,
		db:     db,
		logger: logger,
		signal: make(chan bool),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pruning"
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(block types.Block) error {
	if m.cfg == nil {
		// Nothing to do, pruning is disabled
		return nil
	}

	if block.Height%uint64(m.cfg.GetInterval()) != 0 {
		// Not an interval height, so just skip
		return nil
	}
	m.signal <- true
	return nil
}
