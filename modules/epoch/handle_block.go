package epoch

import "github.com/forbole/soljuno/types"

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.EpochInfo()
	if err != nil {
		return err
	}
	if !m.updateEpoch(info.Epoch) {
		return nil
	}

	return HandleBlock(block)
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

func HandleBlock(block types.Block) error {
	return nil
}
