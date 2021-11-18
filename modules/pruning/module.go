package pruning

import (
	"fmt"

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
}

// NewModule builds a new Module instance
func NewModule(cfg types.PruningConfig, db db.Database, logger logging.Logger) *Module {
	return &Module{
		cfg:    cfg,
		db:     db,
		logger: logger,
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

	pruningDb, ok := m.db.(db.PruningDb)
	if !ok {
		return fmt.Errorf("pruning is enabled, but your database does not implement PruningDb")
	}

	if block.Slot%uint64(m.cfg.GetInterval()) != 0 {
		// Not an interval slot, so just skip
		return nil
	}
	slot := block.Slot - uint64(m.cfg.GetKeepRecent())

	// Prune the blocks before the given slot
	m.logger.Debug("pruning", "module", "pruning", "slot", slot)
	err := pruningDb.Prune(block.Slot - uint64(m.cfg.GetKeepRecent()))
	if err != nil {
		return fmt.Errorf("error while pruning slot %d: %s", slot, err.Error())
	}

	return nil
}
