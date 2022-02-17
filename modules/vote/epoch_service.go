package vote

import (
	"fmt"

	dbtypes "github.com/forbole/soljuno/db/types"
	solanatypes "github.com/forbole/soljuno/solana/types"
)

func (m *Module) ExecEpoch(epoch uint64) error {
	return m.updateValidatorSkipRates(epoch - 1)
}

// updateValidatorSkipRates properly stores the skip rates of all validators inside the database
func (m *Module) updateValidatorSkipRates(epoch uint64) error {
	slots, err := m.db.GetEpochProducedBlocks(epoch)
	if err != nil {
		return err
	}
	if len(slots) == 0 {
		return fmt.Errorf("%d epoch blocks does not exist", epoch)
	}

	endSlot := slots[len(slots)-1]
	schedules, err := m.client.GetLeaderSchedule(epoch * solanatypes.SlotsInEpoch)
	if err != nil {
		return err
	}

	produced := make(map[int]bool)
	for _, slot := range slots {
		produced[int(slot%solanatypes.SlotsInEpoch)] = true
	}

	skipRateRows := make([]dbtypes.ValidatorSkipRateRow, len(schedules))
	count := 0
	end := int(endSlot % solanatypes.SlotsInEpoch)
	for validator, schedule := range schedules {
		total, skip := getSkipRateReference(end, produced, schedule)
		skipRate := float64(skip) / float64(total)
		skipRateRows[count] = dbtypes.NewValidatorSkipRateRow(validator, epoch, skipRate, total, skip)
		count++
	}

	err = m.db.SaveValidatorSkipRates(skipRateRows)
	if err != nil {
		return err
	}
	return m.db.SaveHistoryValidatorSkipRates(skipRateRows)
}

// getSkipRateReference returns the total and skip amount in a epoch of the validator from the given produced map and the validator schedule
func getSkipRateReference(end int, produced map[int]bool, schedule []int) (int, int) {
	var skip int = 0
	var total int = 0
	for _, slotInEpoch := range schedule {
		total++
		if slotInEpoch > end {
			break
		}
		if ok := produced[slotInEpoch]; !ok {
			skip++
		}
	}
	return total, skip
}
