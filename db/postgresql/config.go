package postgresql

import "github.com/forbole/soljuno/db"

var _ db.ConfigDb = &Database{}

// SaveConfigAccount implements the db.ConfigDb
func (db *Database) SaveConfigAccount(address string, slot uint64, owner string, data string) error {
	stmt := `
INSERT INTO config_account
    (address, slot, owner, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    owner = excluded.owner,
	value = excluded.value
WHERE config_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		owner,
		data,
	)
	return err
}
