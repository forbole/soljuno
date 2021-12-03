package epoch

import "github.com/forbole/soljuno/types"

func (m *Module) HandleBlock(block types.Block) error {
	if block.Slot%432000 != 0 {
		return nil
	}
	return nil
}
