package types

import (
	"github.com/forbole/soljuno/types"
)

type TxRow struct {
	Signature       string      `db:"signature"`
	Slot            uint64      `db:"slot"`
	Success         bool        `db:"success"`
	Fee             uint64      `db:"fee"`
	Logs            interface{} `db:"logs"`
	NumInstructions int         `db:"num_instructions"`
	PartitionId     int         `db:"partition_id"`
}

func NewTxRow(signature string, slot uint64, success bool, fee uint64, logs []string, numInstructions int) TxRow {
	return TxRow{
		Signature:       signature,
		Slot:            slot,
		Success:         success,
		Fee:             fee,
		Logs:            logs,
		NumInstructions: numInstructions,
		PartitionId:     int(slot / 1000),
	}
}

func NewTxRowsFromTxs(txs []types.Tx) []TxRow {
	txRows := make([]TxRow, len(txs))
	for i, tx := range txs {
		txRows[i] = NewTxRow(
			tx.Signature,
			tx.Slot,
			tx.Successful(),
			tx.Fee,
			tx.Logs,
			len(tx.Instructions),
		)
	}
	return txRows
}
