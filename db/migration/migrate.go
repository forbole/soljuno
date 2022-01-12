package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	ALTER TABLE multisig RENAME COLUMN m TO minimum; 
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE multisig RENAME COLUMN minimum TO m;
	`)
	return err
}
