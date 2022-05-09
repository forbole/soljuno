package types

import "time"

type BalanceRow struct {
	Address string `db:"address"`
	Slot    uint64 `db:"slot"`
	Balance uint64 `db:"balance"`
}

func NewBalanceRow(address string, slot uint64, balance uint64) BalanceRow {
	return BalanceRow{
		Address: address,
		Slot:    slot,
		Balance: balance,
	}
}

type BalanceHistoryRow struct {
	Address   string    `db:"address"`
	Timestamp time.Time `db:"timestamp"`
	Balance   uint64    `db:"balance"`
}

func NewBalanceHistoryRow(address string, timestamp time.Time, balance uint64) BalanceHistoryRow {
	return BalanceHistoryRow{
		Address:   address,
		Timestamp: timestamp,
		Balance:   balance,
	}
}

func (b BalanceHistoryRow) Equal(other BalanceHistoryRow) bool {
	return b.Address == other.Address &&
		b.Timestamp.Equal(other.Timestamp) &&
		b.Balance == other.Balance
}
