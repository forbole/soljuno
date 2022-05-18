package types

type MissingHeightRow struct {
	Height uint64 `db:"height"`
}

type MissingSlotRangeRow struct {
	Slot uint64 `db:"slot"`
}
