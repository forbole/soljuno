package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	CREATE TABLE token_price_history
	(
		id          TEXT                        NOT NULL,
		price       DECIMAL                     NOT NULL,
		market_cap  BIGINT                      NOT NULL,
		symbol      TEXT                        NOT NULL,
		timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
	);
	CREATE INDEX token_price_history_id_index ON token_price (id);
	CREATE INDEX token_price_history_timestamp_index ON token_price (timestamp);
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	DROP TABLE token_price_history;
	`)
	return err
}
