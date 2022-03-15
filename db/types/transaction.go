package types

import (
	"github.com/forbole/soljuno/types"
)

type TxRow struct {
	Signature       string      `db:"signature"`
	Slot            uint64      `db:"slot"`
	Error           bool        `db:"error"`
	Fee             uint64      `db:"fee"`
	Logs            interface{} `db:"logs"`
	NumInstructions int         `db:"num_instructions"`
	PartitionId     int         `db:"partition_id"`
}

func NewTxRow(signature string, slot uint64, isErr bool, fee uint64, logs []string, numInstructions int) TxRow {
	return TxRow{
		Signature:       signature,
		Slot:            slot,
		Error:           isErr,
		Fee:             fee,
		Logs:            logs,
		NumInstructions: numInstructions,
		PartitionId:     int(slot / 1000),
	}
}

func NewTxRowsFromTxs(txs []types.Tx) ([]TxRow, error) {
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
	return txRows, nil
}
