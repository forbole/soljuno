package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

var _ db.InstructionDb = &Database{}

// SaveInstructions implements db.MsgDb
func (db *Database) SaveInstructions(msgs []dbtypes.InstructionRow) error {
	if len(msgs) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO message
	(transaction_hash, slot, index, inner_index, involved_accounts, program, raw_data, type, value, partition_id) VALUES`
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 10
	params = make([]interface{}, 0, paramsNumber*len(msgs))
	for _, msg := range msgs {
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
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// CreateInstructionsPartition implements db.MsgDb
func (db *Database) CreateInstructionsPartition(id int) error {
	return db.createPartition("message", id)
}

// PruneInstructionsBeforeSlot implements db.MsgDb
func (db *Database) PruneInstructionsBeforeSlot(slot uint64) error {
	for {
		name, err := db.getOldestInstructionsPartitionBeforeSlot(slot)
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

// getOldestInstructionsPartitionBeforeSlot allows to get the oldest msg partition
func (db *Database) getOldestInstructionsPartitionBeforeSlot(slot uint64) (string, error) {
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
