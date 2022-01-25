package txs

import (
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		block := <-m.buffer
		_ = m.pool.Submit(func() {
			txRows, err := dbtypes.NewTxRowsFromTxs(block.Txs)
			if err != nil {
				log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
			}

			err = m.db.SaveTxs(txRows)
			if err != nil {
				log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
			}
		})
	}
}
