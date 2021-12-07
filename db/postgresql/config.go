package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.ConfigDb = &Database{}

// SaveValidatorConfig implements the db.ConfigDb
func (db *Database) SaveValidatorConfig(row dbtypes.ValidatorConfigRow) error {
	stmt := `
INSERT INTO validator_config
    (address, slot, owner, name, keybase_username, website, details)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    owner = excluded.owner,
	name = excluded.name,
	keybase_username = excluded.keybase_username,
	website = excluded.website,
	details = excluded.details
WHERE validator_config.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		row.Address,
		row.Slot,
		row.Owner,
		row.Name,
		row.KeybaseUsername,
		row.Website,
		row.Details,
	)
	return err
}
