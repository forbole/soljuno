package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_price ALTER COLUMN volume TYPE FLOAT;
	ALTER TABLE token_price_history ALTER COLUMN volume TYPE FLOAT;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_price ALTER COLUMN volume TYPE BIGINT;
	ALTER TABLE token_price_history ALTER COLUMN volume TYPE BIGINT;
	`)
	return err
}
