package txs

import (
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		block := <-m.buffer
		_ = m.pool.Submit(func() {
			err := m.db.SaveTxs(block.Txs)
			if err != nil {
				log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
			}
		})
	}
}
