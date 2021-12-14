package epoch

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.EpochInfo()
	if err != nil {
		return err
	}
	if !m.updateEpoch(info.Epoch) {
		return nil
	}
	err = m.db.SaveEpoch(dbtypes.NewEpochInfoRow(m.epoch, info.TransactionCount))
	if err != nil {
		return err
	}
	return handleEpoch(m.epoch, m.db, m.client)
}

func (m *Module) updateEpoch(epoch uint64) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.epoch == epoch {
		return false
	}
	m.epoch = epoch
	return true
}

func handleEpoch(epoch uint64, db db.EpochDb, client client.Proxy) error {
	// NOTE: updateSupplyInfo takes too much time so specificly use goroutine here.
	go func() {
		err := updateSupplyInfo(epoch, db, client)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}()
	err := updateInflationRate(epoch, db, client)
	if err != nil {
		return err
	}
	err = updateEpochScheduleParam(epoch, db, client)
	if err != nil {
		return err
	}
	return updateInflationGovernorParam(epoch, db, client)
}
