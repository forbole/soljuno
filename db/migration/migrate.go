package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_price ADD COLUMN volume BIGINT NOT NULL DEFAULT 0;
	ALTER TABLE token_price_history ADD COLUMN volume BIGINT NOT NULL DEFAULT 0;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_price DROP COLUMN volume;
	ALTER TABLE token_price_history DROP COLUMN volume;
	`)
	return err
}
