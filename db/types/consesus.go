package types

import "time"

// BlockRow represents a single block row stored inside the database
type BlockRow struct {
	Slot      uint64    `db:"slot"`
	Height    uint64    `db:"height"`
	Hash      string    `db:"hash"`
	Proposer  string    `db:"proposer"`
	Timestamp time.Time `db:"timestamp"`
	NumTxs    int       `db:"num_txs"`
}

func NewBlockRow(slot uint64, height uint64, hash string, proposer string, timestamp time.Time, numTxs int) BlockRow {
	return BlockRow{
		Slot:      slot,
		Height:    height,
		Hash:      hash,
		Proposer:  proposer,
		Timestamp: timestamp,
		NumTxs:    numTxs,
	}
}

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
