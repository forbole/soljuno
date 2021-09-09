package postgresql

import "github.com/forbole/soljuno/db"

// type check to ensure interface is properly implemented
var _ db.PruningDb = &Database{}

// GetLastPruned implements db.PruningDb
func (db *Database) GetLastPruned() (uint64, error) {
	var lastPrunedHeight uint64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(last_pruned_slot),0) FROM pruning LIMIT 1;`).Scan(&lastPrunedHeight)
	return lastPrunedHeight, err
}

// StoreLastPruned implements db.PruningDb
func (db *Database) StoreLastPruned(slot uint64) error {
	_, err := db.Sql.Exec(`DELETE FROM pruning`)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`INSERT INTO pruning (last_pruned_slot) VALUES ($1)`, slot)
	return err
}

// Prune implements db.PruningDb
func (db *Database) Prune(slot uint64) error {
	_, err := db.Sql.Exec(`
DELETE FROM message 
USING transaction 
WHERE message.transaction_hash = transaction.hash AND transaction.slot = $1
`, slot)
	return err
}
