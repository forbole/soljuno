package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE block ALTER COLUMN proposer SET NOT NULL;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE block ALTER COLUMN proposer DROP NOT NULL;
	`)
	return err
}
