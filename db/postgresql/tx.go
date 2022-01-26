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
	insertStmt := `INSERT INTO transaction (hash, slot, error, fee, logs, messages, partition_id) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 7
	for i, tx := range txs {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)

		params = append(
			params,
			tx.Hash,
			tx.Slot,
			tx.Error,
			tx.Fee,
			pq.Array(tx.Logs),
			tx.Messages,
			tx.PartitionId,
		)
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

// CreateTxPartition implements db.Database
func (db *Database) CreateTxPartition(Id int) error {
	return db.createPartition("transaction", Id)
}

// PruneMsgsBeforeSlot implements db.MsgDb
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
