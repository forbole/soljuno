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
		for _, s := range m.services {
			err := s.Prune(block.Slot)
			if err != nil {
				log.Error().Str("module", m.Name()).Str("target", s.Name()).Uint64("slot", block.Slot).Err(err).Send()
			}
		}
	}
}
