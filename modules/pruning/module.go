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

	if block.Slot%uint64(m.cfg.GetInterval()) != 0 {
		// Not an interval slot, so just skip
		return nil
	}

	pruningDb, ok := m.db.(db.PruningDb)
	if !ok {
		return fmt.Errorf("pruning is enabled, but your database does not implement PruningDb")
	}

	// Get last pruned slot
	var slot, err = pruningDb.GetLastPruned()
	if err != nil {
		return err
	}

	// Iterate from last pruned slot until (current block slot - keep recent) to
	// avoid pruning the recent blocks that should be kept
	for ; slot < block.Slot-uint64(m.cfg.GetKeepRecent()); slot++ {

		if slot%uint64(m.cfg.GetKeepEvery()) == 0 {
			// The slot should be kept, so just skip
			continue
		}

		// Prune the block by slot
		m.logger.Debug("pruning", "module", "pruning", "slot", slot)
		err = pruningDb.Prune(slot)
		if err != nil {
			return fmt.Errorf("error while pruning slot %d: %s", slot, err.Error())
		}
	}

	return pruningDb.StoreLastPruned(slot)
}
