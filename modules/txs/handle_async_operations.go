package txs

import (
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		block := <-m.buffer
		_, err := m.pool.DoAsync(func() error {
			txRows, err := dbtypes.NewTxRowsFromTxs(block.Txs)
			if err != nil {
				m.handleAsyncError(err, block)
				return nil
			}

			err = m.db.SaveTxs(txRows)
			m.handleAsyncError(err, block)
			return nil
		})
		m.handleAsyncError(err, block)
	}
}

func (m *Module) handleAsyncError(err error, block types.Block) {
	if err != nil {
		log.Error().Str("module", m.Name()).Err(err).Send()
		log.Info().Str("module", m.Name()).Msg("re-enqueueing failed txs")
		m.buffer <- block
	}
}
