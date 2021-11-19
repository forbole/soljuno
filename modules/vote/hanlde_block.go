package vote

import (
	"fmt"

	"github.com/forbole/soljuno/types"
)

func HandleBLock(m *Module, block types.Block) error {
	if m.cfg == nil {
		// Nothing to do, pruning is disabled
		return nil
	}
	if block.Slot%uint64(m.cfg.GetInterval()) != 0 {
		// Not an interval slot, so just skip
		return nil
	}
	slot := block.Slot - uint64(m.cfg.GetKeepRecent())

	// Delete the validator statuses before the given slot
	err := m.db.PruneValidatorStatus(block.Slot - uint64(m.cfg.GetKeepRecent()))
	if err != nil {
		return fmt.Errorf("error while pruning validator statuses %d: %s", slot, err.Error())
	}
	return nil
}
