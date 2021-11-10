package postgresql

import "github.com/forbole/soljuno/db"

// type check to ensure interface is properly implemented
var _ db.BpfLoaderDb = &Database{}

// SaveBufferAccount implements the db.BpfLoaderDb
func (db *Database) SaveBufferAccount(address string, slot uint64, authority string, state string) error {
	stmt := `
INSERT INTO buffer_account
    (address, slot, authority, state)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    authority = excluded.authority,
	state = excluded.state
WHERE buffer_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		authority,
		state,
	)
	return err
}

// DeleteBufferAccount implements the db.BpfLoaderDb
func (db *Database) DeleteBufferAccount(address string) error {
	stmt := `DELETE FROM buffer_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}

// SaveProgramAccount implements the db.BpfLoaderDb
func (db *Database) SaveProgramAccount(address string, slot uint64, programDataAccount string, state string) error {
	stmt := `
INSERT INTO program_account
    (address, slot, program_data_account, state)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    program_data_account = excluded.program_data_account,
	state = excluded.state
WHERE program_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		programDataAccount,
		state,
	)
	return err
}

// DeleteProgramAccount implements the db.BpfLoaderDb
func (db *Database) DeleteProgramAccount(address string) error {
	stmt := `DELETE FROM program_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}

// SaveProgramDataAccount implements the db.BpfLoaderDb
func (db *Database) SaveProgramDataAccount(address string, slot uint64, lastModifiedSlot uint64, updateAuthority string, state string) error {
	stmt := `
INSERT INTO program_data_account
    (address, slot, last_modified_slot, update_authority, state)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    last_modified_slot = excluded.last_modified_slot,
    update_authority = excluded.update_authority,
    state = excluded.state
WHERE program_data_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		lastModifiedSlot,
		updateAuthority,
		state,
	)
	return err
}

// DeleteProgramDataAccount implements the db.BpfLoaderDb
func (db *Database) DeleteProgramDataAccount(address string) error {
	stmt := `DELETE FROM program_data_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}
