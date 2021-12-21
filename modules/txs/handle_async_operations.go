package txs

import (
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		block := <-m.buffer
		err := m.db.SaveTxs(block.Txs)
		go func() {
			if err != nil {
				log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
			}
		}()
	}
}
