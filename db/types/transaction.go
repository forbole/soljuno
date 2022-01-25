package types

import (
	"encoding/json"

	"github.com/forbole/soljuno/types"
)

type TxRow struct {
	Hash        string      `db:"hash"`
	Slot        uint64      `db:"slot"`
	Error       bool        `db:"error"`
	Fee         uint64      `db:"fee"`
	Logs        interface{} `db:"logs"`
	Messages    interface{} `db:"messages"`
	PartitionID int         `db:"partition_id"`
}

func NewTxRow(hash string, slot uint64, isErr bool, fee uint64, logs []string, messages interface{}) TxRow {
	return TxRow{
		Hash:        hash,
		Slot:        slot,
		Error:       isErr,
		Fee:         fee,
		Logs:        logs,
		Messages:    messages,
		PartitionID: int(slot / 1000),
	}
}

func NewTxRowsFromTxs(txs []types.Tx) ([]TxRow, error) {
	txRows := make([]TxRow, len(txs))
	for i, tx := range txs {
		msgs, err := json.Marshal(types.NewSanitizedMessages(tx.Messages))
		if err != nil {
			return nil, err
		}
		txRows[i] = NewTxRow(
			tx.Hash,
			tx.Slot,
			tx.Successful(),
			tx.Fee,
			tx.Logs,
			msgs,
		)
	}
	return txRows, nil
}
