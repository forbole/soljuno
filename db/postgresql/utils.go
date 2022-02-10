package postgresql

import "fmt"

func (db *Database) InsertBatch(insertStmt string, conflictStmt string, params []interface{}, paramsNumber int) error {
	start := 0
	maxParamsAmount := MAX_PARAMS_LENGTH / paramsNumber * paramsNumber
	sliceNumber := len(params) / maxParamsAmount

	for i := 0; i < sliceNumber; i++ {
		paramsStmt := ""
		for j := 0; j < maxParamsAmount; j += paramsNumber {
			paramsStmt += getParamsStmt(j, paramsNumber)
		}
		err := db.insertWithParams(insertStmt, paramsStmt, conflictStmt, params[start:start+maxParamsAmount])
		if err != nil {
			return err
		}
		start += maxParamsAmount
	}

	// store the rest of params
	paramsStmt := ""
	for curr := 0; curr < len(params)-start; curr += paramsNumber {
		paramsStmt += getParamsStmt(curr, paramsNumber)
	}
	return db.insertWithParams(insertStmt, paramsStmt, conflictStmt, params[start:])
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
