package postgresql

import (
	"database/sql"
	"fmt"
)

func (db *Database) InsertBatch(insertStmt string, conflictStmt string, params []interface{}, paramsNumber int) error {
	start := 0
	maxParamsAmount := MAX_PARAMS_LENGTH / paramsNumber * paramsNumber
	sliceNumber := len(params) / maxParamsAmount
	errChs := make(chan error)

	for i := 0; i <= sliceNumber; i++ {
		paramsStmt := ""
		var amount = maxParamsAmount
		// the last run
		if i == sliceNumber {
			amount = len(params) - start
		}

		for j := 0; j < amount; j += paramsNumber {
			paramsStmt += getParamsStmt(j, paramsNumber)
		}
		go func(start int, end int) {
			errChs <- db.insertWithParams(insertStmt, paramsStmt, conflictStmt, params[start:end])
		}(start, start+amount)
		start += amount
	}

	for i := 0; i <= sliceNumber; i++ {
		err := <-errChs
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) insertWithParams(insertStmt string, paramsStmt string, conflictStmt string, params []interface{}) error {
	if len(params) == 0 {
		return nil
	}
	paramsStmt = paramsStmt[:len(paramsStmt)-1]
	insertStmt += paramsStmt
	insertStmt += conflictStmt
	_, err := db.Sqlx.Exec(insertStmt, params...)
	return err
}

func getParamsStmt(start, number int) string {
	stmt := "("
	for i := 1; i <= number; i++ {
		stmt += fmt.Sprintf("$%d,", start+i)
	}
	stmt = stmt[:len(stmt)-1]
	return stmt + "),"
}

// createPartition allows to create a partition with the id for the given table name
func (db *Database) createPartition(table string, id int) error {
	stmt := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v_%d PARTITION OF %v FOR VALUES IN (%d)",
		table,
		id,
		table,
		id,
	)
	_, err := db.Exec(stmt)
	return err
}

// getOldestPartitionBeforeSlot allows to get the oldest tx partition
func (db *Database) getOldestPartitionBeforeSlot(name string, slot uint64) (string, error) {
	stmt := fmt.Sprintf(`
	SELECT tableoid::pg_catalog.regclass FROM %s WHERE slot <= $1 ORDER BY slot ASC LIMIT 1;
	`, name)
	var partitionName string
	err := db.Sqlx.QueryRow(stmt, slot).Scan(&partitionName)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return partitionName, nil
}

// dropPartition allows to drop a partition with the given partition name
func (db *Database) dropPartition(name string) error {
	stmt := fmt.Sprintf(
		"DROP TABLE IF EXISTS %v",
		name,
	)
	_, err := db.Exec(stmt)
	return err
}
