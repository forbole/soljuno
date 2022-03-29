package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

// type check to ensure interface is properly implemented
var _ db.BpfLoaderDb = &Database{}

// SaveBufferAccount implements the db.BpfLoaderDb
func (db *Database) SaveBufferAccount(account dbtypes.BufferAccountRow) error {
	stmt := `
INSERT INTO buffer_account
    (address, slot, authority)
VALUES ($1, $2, $3)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    authority = excluded.authority
WHERE buffer_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		account.Address,
		account.Slot,
		account.Authority,
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
func (db *Database) SaveProgramAccount(account dbtypes.ProgramAccountRow) error {
	stmt := `
INSERT INTO program_account
    (address, slot, program_data_account)
VALUES ($1, $2, $3)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    program_data_account = excluded.program_data_account
WHERE program_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		account.Address,
		account.Slot,
		account.ProgramDataAccount,
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
func (db *Database) SaveProgramDataAccount(account dbtypes.ProgramDataAccountRow) error {
	stmt := `
INSERT INTO program_data_account
    (address, slot, last_modified_slot, update_authority)
VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    last_modified_slot = excluded.last_modified_slot,
    update_authority = excluded.update_authority
WHERE program_data_account.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		account.Address,
		account.Slot,
		account.LastModifiedSlot,
		account.UpdateAuthority,
	)
	return err
}

// DeleteProgramDataAccount implements the db.BpfLoaderDb
func (db *Database) DeleteProgramDataAccount(address string) error {
	stmt := `DELETE FROM program_data_account WHERE address = $1`
	_, err := db.Sqlx.Exec(stmt, address)
	return err
}
