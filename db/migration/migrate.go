package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		ALTER TABLE token_delegation DROP CONSTRAINT token_delegation_delegate_address_fk;
		DROP INDEX token_delegation_delegate_address_index;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_delegation ADD CONSTRAINT token_delegation_delegate_address_fk
	FOREIGN KEY (delegate_address) REFERENCES token_account(address) ON DELETE CASCADE;
	CREATE INDEX token_delegation_delegate_address_index ON token_delegation (delegate_address);
	`)
	return err
}
