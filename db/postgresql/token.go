package postgresql

import (
	"strconv"

	"github.com/forbole/soljuno/db"
	"github.com/lib/pq"
)

var _ db.TokenDb = &Database{}

// SaveToken implements the db.TokenDb
func (db *Database) SaveToken(
	address string,
	slot uint64,
	decimals uint8,
	mintAuthority string,
	freezeAuthority string,
) error {
	stmt := `
INSERT INTO token
    (address, slot, decimals, mint_authority, freeze_authority)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
	decimals = excluded.decimals,
	mint_authority = excluded.mint_authority,
	freeze_authority = excluded.freeze_authority
WHERE token.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		decimals,
		mintAuthority,
		freezeAuthority,
	)
	return err
}

// SaveTokenAccount implements the db.TokenDb
func (db *Database) SaveTokenAccount(address string, slot uint64, token, owner string) error {
	stmt := `
INSERT INTO token_account
	(address, slot, token, owner)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	token = excluded.token,
	owner = excluded.owner
WHERE token_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		token,
		owner,
	)
	return err
}

// DeleteTokenAccount implements the db.TokenDb
func (db *Database) DeleteTokenAccount(address string) error {
	stmt := `DELETE FROM token_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}

// SaveMultisig implements the db.TokenDb
func (db *Database) SaveMultisig(address string, slot uint64, singers []string, m uint8) error {
	stmt := `
INSERT INTO multisig
	(address, slot, signers, m)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	signers = excluded.signers,
	m = excluded.m
WHERE multisig.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		pq.Array(singers),
		m,
	)
	return err
}

// SaveTokenDelegate implements the db.TokenDb
func (db *Database) SaveTokenDelegation(source string, delegate string, slot uint64, amount uint64) error {
	stmt := `
INSERT INTO token_delegation
	(source_address, delegate_address, slot, amount)
VALUES ($1, $2, $3, $4)
ON CONFLICT (source_address) DO UPDATE
	SET slot = excluded.slot,
	amount = excluded.amount
WHERE token_delegation.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		source,
		delegate,
		slot,
		amount,
	)
	return err
}

// DeleteTokenDelegation implements the db.TokenDb
func (db *Database) DeleteTokenDelegation(address string) error {
	stmt := `DELETE FROM token_delegation WHERE source_address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}

// SaveTokenSupply implements the db.TokenDb
func (db *Database) SaveTokenSupply(address string, slot uint64, supply uint64) error {
	stmt := `
INSERT INTO token_supply
	(address, slot, supply)
VALUES ($1, $2, $3)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	supply = excluded.supply
WHERE token_supply.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		strconv.FormatUint(supply, 10),
	)
	return err
}
