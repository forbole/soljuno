package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

var _ db.InstructionDb = &Database{}

// SaveInstructions implements db.InstructionDb
func (db *Database) SaveInstructions(instructions []dbtypes.InstructionRow) error {
	if len(instructions) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO instruction
	(tx_signature, slot, index, inner_index, involved_accounts, program, raw_data, type, value, partition_id) VALUES`
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 10
	params = make([]interface{}, 0, paramsNumber*len(instructions))
	for _, instruction := range instructions {
		params = append(
			params,
			instruction.TxSignature,
			instruction.Slot,
			instruction.Index,
			instruction.InnerIndex,
			pq.Array(instruction.InvolvedAccounts),
			instruction.Program,
			instruction.RawData,
			instruction.Type,
			instruction.Value,
			instruction.PartitionId,
		)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// CreateInstructionsPartition implements db.InstructionDb
func (db *Database) CreateInstructionPartition(id int) error {
	return db.createPartition("instruction", id)
}

// PruneInstructionsBeforeSlot implements db.InstructionDb
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

// getOldestInstructionsPartitionBeforeSlot allows to get the oldest  partition
func (db *Database) getOldestInstructionsPartitionBeforeSlot(slot uint64) (string, error) {
	stmt := `
	SELECT tableoid::pg_catalog.regclass FROM instruction WHERE slot <= $1 ORDER BY slot ASC LIMIT 1;
	`
	var partitionName string
	err := db.Sqlx.QueryRow(stmt, slot).Scan(&partitionName)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return partitionName, nil
}
