package postgresql

import (
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

// SaveMessages implements db.Database
func (db *Database) SaveMessages(msgs []dbtypes.MsgRow) error {
	if len(msgs) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO message
	(transaction_hash, slot, index, inner_index, involved_accounts, program, raw_data, type, value, partition_id) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 10
	for i, msg := range msgs {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(
			params,
			msg.TxHash,
			msg.Slot,
			msg.Index,
			msg.InnerIndex,
			pq.Array(msg.InvolvedAccounts),
			msg.Program,
			msg.RawData,
			msg.Type,
			msg.Value,
			msg.PartitionId,
		)
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}
