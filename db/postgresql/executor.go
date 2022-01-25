package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
)

var _ db.ExcecutorDb = &Database{}

func (db *Database) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return db.Sqlx.Exec(sql, args...)
}

func (db *Database) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return db.Sqlx.Query(sql, args...)
}
