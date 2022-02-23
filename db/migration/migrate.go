package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE transaction DROP COLUMN messages;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE transaction ADD COLUMN messages JSON;
	`)
	return err
}
