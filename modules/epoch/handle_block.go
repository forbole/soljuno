package epoch

import (
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.GetEpochInfo()
	if err != nil {
		return err
	}
	if !m.updateEpoch(info.Epoch) {
		return nil
	}
	return m.handleEpoch()
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

func (m *Module) handleEpoch() error {
	// NOTE: updateSupplyInfo takes too much time so specificly use goroutine here.
	go func() {
		err := updateSupplyInfo(m.epoch, m.db, m.client)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}()
	return m.RunServices()
}
