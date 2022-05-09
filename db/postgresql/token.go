package postgresql

import (
	"strconv"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"

	"github.com/lib/pq"
)

var _ db.TokenDb = &Database{}

// SaveToken implements the db.TokenDb
func (db *Database) SaveToken(
	token dbtypes.TokenRow,
) error {
	stmt := `
INSERT INTO token
    (mint, slot, decimals, mint_authority, freeze_authority)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (mint) DO UPDATE
    SET slot = excluded.slot,
	decimals = excluded.decimals,
	mint_authority = excluded.mint_authority,
	freeze_authority = excluded.freeze_authority
WHERE token.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		token.Mint,
		token.Slot,
		token.Decimals,
		token.MintAuthority,
		token.FreezeAuthority,
	)
	return err
}

// SaveTokenAccount implements the db.TokenDb
func (db *Database) SaveTokenAccount(account dbtypes.TokenAccountRow) error {
	stmt := `
INSERT INTO token_account
	(address, slot, mint, owner)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	mint = excluded.mint,
	owner = excluded.owner
WHERE token_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		account.Address,
		account.Slot,
		account.Mint,
		account.Owner,
	)
	return err
}

// DeleteTokenAccount implements the db.TokenDb
func (db *Database) DeleteTokenAccount(address string) error {
	_, err := db.Sqlx.Exec(`DELETE FROM token_account WHERE address = $1`, address)
	if err != nil {
		return err
	}
	_, err = db.Sqlx.Exec(`DELETE FROM token_account_balance WHERE address =$1`, address)
	return err
}

// SaveMultisig implements the db.TokenDb
func (db *Database) SaveMultisig(multisig dbtypes.MultisigRow) error {
	stmt := `
INSERT INTO multisig
	(address, slot, signers, minimum)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	signers = excluded.signers,
	minimum = excluded.minimum
WHERE multisig.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		multisig.Address,
		multisig.Slot,
		pq.Array(multisig.Signers),
		multisig.Minimum,
	)
	return err
}

// SaveTokenDelegate implements the db.TokenDb
func (db *Database) SaveTokenDelegation(delegation dbtypes.TokenDelegationRow) error {
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
		delegation.Source,
		delegation.Destination,
		delegation.Slot,
		strconv.FormatUint(delegation.Amount, 10),
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
func (db *Database) SaveTokenSupply(supply dbtypes.TokenSupplyRow) error {
	stmt := `
INSERT INTO token_supply
	(mint, slot, supply)
VALUES ($1, $2, $3)
ON CONFLICT (mint) DO UPDATE
	SET slot = excluded.slot,
	supply = excluded.supply
WHERE token_supply.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		supply.Mint,
		supply.Slot,
		strconv.FormatUint(supply.Supply, 10),
	)
	return err
}
