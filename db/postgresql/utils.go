package postgresql

import "fmt"

func (db *Database) insertWithParams(insertStmt string, paramsStmt string, conflictStmt string, params []interface{}) error {
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
