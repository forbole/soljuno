package epoch

import "github.com/forbole/soljuno/types"

func (m *Module) HandleBlock(block types.Block) error {
	info, err := m.client.EpochInfo()
	if err != nil {
		return err
	}
	if info.Epoch == m.epoch {
		return nil
	}
	return nil
}
