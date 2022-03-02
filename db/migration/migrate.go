package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX message_slot_index RENAME TO message_slot_old_index;
	CREATE INDEX message_slot_index ON message (slot DESC);
	DROP INDEX message_slot_old_index;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX message_slot_index RENAME TO message_slot_old_index;
	CREATE INDEX message_slot_index ON message (slot);
	DROP INDEX message_slot_old_index;
	`)
	return err
}
