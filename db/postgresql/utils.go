package postgresql

import "fmt"

func (db *Database) InsertBatch(insertStmt string, conflictStmt string, params []interface{}, paramsNumber int) error {
	start := 0
	maxParamsAmount := MAX_PARAMS_LENGTH / paramsNumber * paramsNumber
	sliceNumber := len(params) / maxParamsAmount
	errChs := make(chan error)

	fmt.Println(len(params))

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
