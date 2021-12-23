package pruning

import (
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		<-m.signal
		block, err := m.db.GetLastBlock()
		if err != nil {
			log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
		}
		pruningSlot := block.Slot - uint64(m.cfg.GetKeepRecent())
		for _, s := range m.services {
			err := s.Prune(pruningSlot)
			if err != nil {
				log.Error().Str("module", m.Name()).Str("target", s.Name()).Uint64("slot", pruningSlot).Err(err).Send()
			}
		}
	}
}
