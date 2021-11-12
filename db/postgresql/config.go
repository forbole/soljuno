package postgresql

import (
	"encoding/json"

	"github.com/forbole/soljuno/db"
)

var _ db.ConfigDb = &Database{}

type parsedInfo struct {
	Name            string `json:"name"`
	KeybaseUsername string `json:"keybaseUsername"`
	Website         string `json:"website"`
	Details         string `json:"details"`
}

// SaveConfigAccount implements the db.ConfigDb
func (db *Database) SaveConfigAccount(address string, slot uint64, owner string, data string) error {
	var parsed parsedInfo
	err := json.Unmarshal([]byte(data), &parsed)
	if err != nil {
		return err
	}

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
	_, err = db.Sqlx.Exec(
		stmt,
		address,
		slot,
		owner,
		parsed.Name,
		parsed.KeybaseUsername,
		parsed.Website,
		parsed.Details,
	)
	return err
}
