package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

var _ db.MsgDb = &Database{}

// SaveMessages implements db.MsgDb
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

// CreateMsgPartition implements db.MsgDb
func (db *Database) CreateMsgPartition(id int) error {
	return db.createPartition("message", id)
}

// PruneMsgsBeforeSlot implements db.MsgDb
func (db *Database) PruneMsgsBeforeSlot(slot uint64) error {
	for {
		name, err := db.getOldestMsgPartitionBeforeSlot(slot)
		if err != nil {
			return err
		}
		if name == "" {
			return nil
		}

		err = db.dropPartition(name)
		if err != nil {
			return err
		}
	}
}

// getOldestMsgPartitionBeforeSlot allows to get the oldest msg partition
func (db *Database) getOldestMsgPartitionBeforeSlot(slot uint64) (string, error) {
	stmt := `
	SELECT tableoid::pg_catalog.regclass FROM message WHERE slot <= $1 ORDER BY slot ASC LIMIT 1;
	`
	var partitionName string
	err := db.Sqlx.QueryRow(stmt, slot).Scan(&partitionName)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return partitionName, nil
}
