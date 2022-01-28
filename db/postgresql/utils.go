package postgresql

import "fmt"

func (db *Database) InsertBatch(insertStmt string, conflictStmt string, params []interface{}, paramsNumber int) error {
	sliceNumber := len(params) / MAX_PARAMS_LENGTH
	start := 0
	// if slice is larger than MAX_PARAMS_LENGTH
	for i := 0; i < sliceNumber; i += sliceNumber {
		paramsStmt := ""
		for j := 0; j < MAX_PARAMS_LENGTH; j += paramsNumber {
			paramsStmt += getParamsStmt(j, paramsNumber)
		}

		end := (i + 1) * MAX_PARAMS_LENGTH / paramsNumber * paramsNumber
		err := db.insertWithParams(insertStmt, paramsStmt, conflictStmt, params[start:end])
		if err != nil {
			return err
		}
		start = end
	}

	// store the rest of params
	paramsStmt := ""
	for curr := 0; curr < len(params)%MAX_PARAMS_LENGTH; curr += paramsNumber {
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
