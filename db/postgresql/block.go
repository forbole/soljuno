package postgresql

import (
	"database/sql"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.BlockDb = &Database{}

// SaveBlock implements db.BlockDb
func (db *Database) SaveBlock(block dbtypes.BlockRow) error {
	stmt := `
INSERT INTO block (slot, height, hash, leader, timestamp, num_txs)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`
	leader := sql.NullString{Valid: len(block.Leader) != 0, String: block.Leader}
	_, err := db.Sqlx.Exec(
		stmt, block.Slot, block.Height, block.Hash, leader, block.Timestamp, block.NumTxs,
	)
	return err
}

// HasBlock implements db.BlockDb
func (db *Database) HasBlock(height uint64) (bool, error) {
	var res bool
	err := db.Sqlx.QueryRow(`SELECT EXISTS(SELECT 1 FROM block WHERE slot = $1);`, height).Scan(&res)
	return res, err
}
