package fix

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

func HandleBlock(block types.Block, db db.FixMissingBlockDb, queue types.SlotQueue, client ClientProxy) error {
	slot := block.Slot - block.Slot%100
	// fix missing latest slot
	go EnqueueMissingSlots(db, queue, client, slot-100, slot)
	return nil
}

func EnqueueMissingSlots(
	db db.FixMissingBlockDb,
	queue types.SlotQueue,
	client ClientProxy,
	start uint64,
	end uint64,
) {
	for i := start; i < end; {
		next := end - 1
		height, err := db.GetMissingHeight(i, next)
		if err != nil {
			continue
		}
		// Skip if height = 0 meaning that the given range is no missing blocks there
		if height == 0 {
			i = next + 1
			continue
		}

		rangeStart, rangeEnd, err := db.GetMissingSlotRange(height)
		if err != nil {
			continue
		}
		// Skip if end = 0 meaning that the given height is not missing
		if rangeEnd == 0 {
			i = next + 1
			continue
		}
		slots, err := client.GetBlocks(rangeStart, rangeEnd)
		if err != nil {
			continue
		}

		// The slots must be larger than 0 since the height is missing
		if len(slots) == 0 {
			continue
		}
		for _, slot := range slots {
			queue <- slot
		}

		i = rangeEnd + 1
	}
}
