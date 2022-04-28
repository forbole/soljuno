package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

var _ db.TxDb = &Database{}

// SaveTxs implements db.Database
func (db *Database) SaveTxs(txs []dbtypes.TxRow) error {
	if len(txs) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO transaction (signature, slot, index, involved_accounts, success, fee, logs, num_instructions, partition_id) VALUES`
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 8
	params = make([]interface{}, 0, paramsNumber*len(txs))
	for _, tx := range txs {
		params = append(
			params,
			tx.Signature,
			tx.Slot,
			tx.Index,
			tx.InvolvedAccounts,
			tx.Success,
			tx.Fee,
			pq.Array(tx.Logs),
			tx.NumInstructions,
			tx.PartitionId,
		)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// CreateTxPartition implements db.Database
func (db *Database) CreateTxPartition(Id int) error {
	return db.createPartition("transaction", Id)
}

// PruneTxsBeforeSlot implements db.TxDb
func (db *Database) PruneTxsBeforeSlot(slot uint64) error {
	for {
		name, err := db.getOldestTxPartitionBeforeSlot(slot)
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

// getOldestTxPartitionBeforeSlot allows to get the oldest tx partition
func (db *Database) getOldestTxPartitionBeforeSlot(slot uint64) (string, error) {
	stmt := `
	SELECT tableoid::pg_catalog.regclass FROM transaction WHERE slot <= $1 ORDER BY slot ASC LIMIT 1;
	`
	var partitionName string
	err := db.Sqlx.QueryRow(stmt, slot).Scan(&partitionName)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return partitionName, nil
}
