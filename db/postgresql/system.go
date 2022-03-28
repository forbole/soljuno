package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.SystemDb = &Database{}

// SaveNonceAccount implements the db.SystemDb
func (db *Database) SaveNonceAccount(
	account dbtypes.NonceAccountRow,
) error {
	stmt := `
INSERT INTO nonce_account
	(address, slot, authority, blockhash, lamports_per_signature)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	authority = excluded.authority,
	blockhash = excluded.blockhash,
	lamports_per_signature = excluded.lamports_per_signature
WHERE nonce_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		account.Address,
		account.Slot,
		account.Authority,
		account.Blockhash,
		account.LamportsPerSignature,
	)
	return err
}

// DeleteNonceAccount implements the db.SystemDb
func (db *Database) DeleteNonceAccount(address string) error {
	stmt := `DELETE FROM nonce_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}
