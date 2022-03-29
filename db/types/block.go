package types

import "time"

// BlockRow represents a single block row stored inside the database
type BlockRow struct {
	Slot      uint64    `db:"slot"`
	Height    uint64    `db:"height"`
	Hash      string    `db:"hash"`
	Leader    string    `db:"leader"`
	Timestamp time.Time `db:"timestamp"`
	NumTxs    int       `db:"num_txs"`
}

func NewBlockRow(slot uint64, height uint64, hash string, leader string, timestamp time.Time, numTxs int) BlockRow {
	return BlockRow{
		Slot:      slot,
		Height:    height,
		Hash:      hash,
		Leader:    leader,
		Timestamp: timestamp,
		NumTxs:    numTxs,
	}
}

func (b BlockRow) Equal(otherBlock BlockRow) bool {
	return b.Slot == otherBlock.Slot &&
		b.Height == otherBlock.Height &&
		b.Hash == otherBlock.Hash &&
		b.Leader == otherBlock.Leader &&
		b.Timestamp.Equal(otherBlock.Timestamp) &&
		b.NumTxs == otherBlock.NumTxs
}
