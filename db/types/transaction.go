package types

import (
	"github.com/forbole/soljuno/types"
)

type TxRow struct {
	Hash        string      `db:"hash"`
	Slot        uint64      `db:"slot"`
	Error       bool        `db:"error"`
	Fee         uint64      `db:"fee"`
	Logs        interface{} `db:"logs"`
	PartitionId int         `db:"partition_id"`
}

func NewTxRow(hash string, slot uint64, isErr bool, fee uint64, logs []string) TxRow {
	return TxRow{
		Hash:        hash,
		Slot:        slot,
		Error:       isErr,
		Fee:         fee,
		Logs:        logs,
		PartitionId: int(slot / 1000),
	}
}

func NewTxRowsFromTxs(txs []types.Tx) ([]TxRow, error) {
	txRows := make([]TxRow, len(txs))
	for i, tx := range txs {
		txRows[i] = NewTxRow(
			tx.Hash,
			tx.Slot,
			tx.Successful(),
			tx.Fee,
			tx.Logs,
		)
	}
	return txRows, nil
}
