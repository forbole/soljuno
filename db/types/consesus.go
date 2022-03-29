package types

// AverageTimeRow is the average slot time each minute/hour/day
type AverageTimeRow struct {
	OneRowID    bool    `db:"one_row_id"`
	Slot        uint64  `db:"slot"`
	AverageTime float64 `db:"average_time"`
}

func NewAverageTimeRow(slot uint64, averageTime float64) AverageTimeRow {
	return AverageTimeRow{
		OneRowID:    true,
		Slot:        slot,
		AverageTime: averageTime,
	}
}
