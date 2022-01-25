package postgresql

import (
	"database/sql"
	"fmt"

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
			tx.PartitionID,
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
func (db *Database) CreateTxPartition(ID int) error {
	stmt := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS transaction_%d PARTITION OF transaction FOR VALUES IN (%d)",
		ID,
		ID,
	)
	_, err := db.Exec(stmt)
	return err
}

// DropTxPartition implements db.Database
func (db *Database) DropTxPartition(name string) error {
	stmt := fmt.Sprintf(
		"DROP TABLE IF EXISTS %v",
		name,
	)
	_, err := db.Exec(stmt)
	return err
}

// GetOldestTxPartitionNameBySlot implements db.Database
func (db *Database) GetOldestTxPartitionNameBeforeSlot(slot uint64) (string, error) {
	stmt := `
	SELECT tableoid::pg_catalog.regclassFROM transaction WHERE slot <= $1 ORDER BY slot ASC LIMIT 1;
	`
	var partitionName string
	err := db.Sqlx.QueryRow(stmt, slot).Scan(&partitionName)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return partitionName, nil
}
