package epoch

import "github.com/forbole/soljuno/types"

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.EpochInfo()
	if err != nil {
		return err
	}
	m.updateEpoch(info.Epoch)

	return nil
}

func (m *Module) updateEpoch(epoch uint64) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.epoch == epoch {
		return
	}
	m.epoch = epoch
}
