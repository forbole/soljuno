package postgresql

import "github.com/forbole/soljuno/db"

// type check to ensure interface is properly implemented
var _ db.PruningDb = &Database{}

// Prune implements db.PruningDb
func (db *Database) Prune(slot uint64) error {
	_, err := db.Sqlx.Exec(`
DELETE FROM message 
WHERE slot <= $1
`, slot)
	if err != nil {
		return err
	}
	_, err = db.Sqlx.Exec(`
DELETE FROM transaction 
WHERE slot <= $1
`, slot)
	return err
}
