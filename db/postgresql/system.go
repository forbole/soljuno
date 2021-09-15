package postgresql

import (
	solana "github.com/forbole/soljuno/solana/types"
)

// SaveToken implements the db.SystemDb
func (db *Database) SaveNonce(
	address string,
	slot uint64,
	authority string,
	blockhash string,
	feeCalculator solana.FeeCalculator,
	state string,
) error {
	stmt := `
    INSERT INTO nonce
        (address, slot, authority, blockhash, lamports_per_signature, state)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT (mint) DO UPDATE
        SET slot = excluded.slot,
        authority = excluded.authority,
        blockhash = excluded.blockhash,
        lamports_per_signature = excluded.lamports_per_signature,
        state = excluded.state
    WHERE nonce.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		authority,
		blockhash,
		feeCalculator.LamportsPerSignature,
		state,
	)
	return err
}
