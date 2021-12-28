package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	DROP TABLE token_unit;
	DROP TABLE token_price;
	CREATE TABLE token_unit
	(
		address     TEXT    NOT NULL PRIMARY KEY,
		price_id    TEXT    NOT NULL DEFAULT '',
		unit_name   TEXT    NOT NULL DEFAULT '',
		logo_uri    TEXT    NOT NULL DEFAULT '',
		description TEXT    NOT NULL DEFAULT '',
		website     TEXT    NOT NULL DEFAULT ''
	);
	CREATE INDEX token_unit_price_id_index ON token_unit (price_id);

	CREATE TABLE token_price
	(
		id          TEXT                        NOT NULL PRIMARY KEY,
		price       DECIMAL                     NOT NULL,
		market_cap  BIGINT                      NOT NULL,
		symbol      TEXT                        NOT NULL,
		timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
	);
	CREATE INDEX token_price_timestamp_index ON token_price (timestamp);
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	DROP TABLE token_unit;
	DROP TABLE token_price;
	CREATE TABLE token_unit
	(
		price_id    TEXT    PRIMARY KEY,
		address     TEXT    NOT NULL UNIQUE,
		unit_name   TEXT    NOT NULL UNIQUE
	);
	CREATE INDEX token_unit_price_id_index ON token_unit (price_id);

	CREATE TABLE token_price
	(
    	unit_name   TEXT                        NOT NULL REFERENCES token_unit (unit_name) PRIMARY KEY,
		price       DECIMAL                     NOT NULL,
   		market_cap  BIGINT                      NOT NULL,
		timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
	)
	CREATE INDEX token_price_timestamp_index ON token_price (timestamp);
	
	`)
	return err
}
