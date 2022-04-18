package txs

import (
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		m.HandleBuffer()
	}
}

func (m *Module) HandleBuffer() {
	block := <-m.buffer
	_, err := m.pool.DoAsync(func() error {
		txRows := dbtypes.NewTxRowsFromTxs(block.Txs)
		err := m.db.SaveTxs(txRows)
		m.HandleAsyncError(err, block)
		return nil
	})
	m.HandleAsyncError(err, block)
}

func (m *Module) HandleAsyncError(err error, block types.Block) {
	if err != nil {
		log.Error().Str("module", m.Name()).Err(err).Send()
		log.Info().Str("module", m.Name()).Msg("re-enqueueing failed txs")
		m.buffer <- block
	}
}
