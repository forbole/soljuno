package txs

import (
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		block := <-m.buffer
		err := <-m.pool.DoAsync(func() error {
			txRows, err := dbtypes.NewTxRowsFromTxs(block.Txs)
			if err != nil {
				return err
			}
			err = m.db.SaveTxs(txRows)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Error().Str("module", m.Name()).Err(err).Send()
			log.Info().Str("module", m.Name()).Msg("re-enqueueing failed txs")
			m.buffer <- block
		}
	}
}
