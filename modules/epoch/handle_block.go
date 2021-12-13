package epoch

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/types"
)

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.EpochInfo()
	if err != nil {
		return err
	}
	if !m.updateEpoch(info.Epoch) {
		return nil
	}

	return handleEpoch(block.Slot, m.client)
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

func handleEpoch(slot uint64, client client.Proxy) error {
	client.InflationRate()
	client.EpochSchedule()
	client.Supply()
	return nil
}
