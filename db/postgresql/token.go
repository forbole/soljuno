package postgresql

import (
	"github.com/forbole/soljuno/db"
	"github.com/lib/pq"
)

var _ db.TokenDb = &Database{}

// SaveToken implements the db.Token
func (db *Database) SaveToken(
	mint string,
	slot uint64,
	decimals uint8,
	mintAuthority string,
	freezeAuthority string,
) error {
	stmt := `
INSERT INTO token
    (mint, slot, decimals, mint_authority, freeze_authority)
VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT (mint) DO UPDATE
    SET slot = excluded.slot,
	decimals = excluded.decimals
	mint_authority = excluded.mint_authority
	freeze_authority = excluded.freeze_authority
WHERE token.slot <= excluded.slot`
	_, err := db.Sql.Exec(
		stmt,
		mint,
		slot,
		decimals,
		mintAuthority,
		freezeAuthority,
	)
	return err
}

// SaveTokenAccount implements the db.Token
func (db *Database) SaveTokenAccount(address string, slot uint64, mint, owner string) error {
	stmt := `
INSERT INTO token_account
	(address, slot, mint, owner)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address)
	SET slot = excluded.slot
	mint = excluded.mint
	owner = excluded.owner
WHERE token_account.slot <= excluded.slot`
	_, err := db.Sql.Exec(
		stmt,
		address,
		slot,
		mint,
		owner,
	)
	return err
}

// SaveMultisig implements the db.Token
func (db *Database) SaveMultisig(address string, slot uint64, singers []string, m uint8) error {
	stmt := `
INSERT INTO multisig
	(address, slot, signers, m)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address)
	SET slot = excluded.slot
	signers = excluded.signers
	m = excluded.m
WHERE token_account.slot <= excluded.slot`
	_, err := db.Sql.Exec(
		stmt,
		address,
		slot,
		pq.Array(singers),
		m,
	)
	return err
}
