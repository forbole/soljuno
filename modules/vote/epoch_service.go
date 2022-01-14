package vote

import (
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
	for validator, schedule := range schedules {
		skipRate := CalculateSkipRate(int(endSlot%solanatypes.SlotsInEpoch), produced, schedule)
		skipRateRows[count] = dbtypes.NewValidatorSkipRateRow(validator, epoch, skipRate)
	}

	return m.db.SaveValidatorSkipRates(skipRateRows)
}

// CalculateSkipRate returns the skip rate of the validator from the given produced map and the validator schedule
func CalculateSkipRate(end int, produced map[int]bool, schedule []int) float64 {
	count := 0
	producedCount := 0
	for _, slotInEpoch := range schedule {
		if slotInEpoch > end {
			break
		}
		count++
		if ok := produced[slotInEpoch]; !ok {
			producedCount++
		}
	}
	return 1 - float64(producedCount/count)
}
