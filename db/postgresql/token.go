package postgresql

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
(mint) DO UPDATE
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

func (db *Database) InitializeAccount()
