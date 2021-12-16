package pruning

import (
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		<-m.signal
		m.prune()
	}
}

func (m *Module) prune() {
	block, err := m.db.GetLastBlock()
	if err != nil {
		log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
	}
	slot := block.Slot - uint64(m.cfg.GetKeepRecent())
	// Prune the blocks before the given slot
	m.logger.Debug("pruning", "module", "pruning", "slot", slot)
	err = m.db.Prune(block.Slot - uint64(m.cfg.GetKeepRecent()))
	if err != nil {
		log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()

	}
}
