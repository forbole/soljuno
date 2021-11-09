package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
)

var _ db.ExceutorDb = &Database{}

func (db *Database) Exec(sql string) (sql.Result, error) {
	return db.Sqlx.Exec(sql)
}
