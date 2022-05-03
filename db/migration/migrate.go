package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		CREATE TABLE transaction_by_address
		(
			address         TEXT    NOT NULL,
			slot            BIGINT  NOT NULL, 
			signature       TEXT    NOT NULL,
			index           INT     NOT NULL DEFAULT 0,
			partition_id    INT     NOT NULL
		) PARTITION BY LIST(partition_id);
		ALTER TABLE transaction_by_address ADD UNIQUE (address, signature, partition_id);
		CREATE INDEX transaction_by_address_slot_index ON transaction_by_address(slot);
		CREATE INDEX transaction_by_address_signature_index ON transaction_by_address(signature);
		CREATE INDEX transaction_by_address_search_index ON transaction_by_address (address, slot DESC, index DESC);

		CREATE OR REPLACE FUNCTION transactions_by_address_internal(
			"target"    TEXT,
			"current"   TEXT = '',
			"limit"     INT = 10
			)
			RETURNS SETOF transaction AS
		$$
		BEGIN
			IF "current" = '' THEN
				RETURN QUERY SELECT 
					t.*
				FROM (
					SELECT signature, partition_id FROM transaction_by_address WHERE address = "target" ORDER BY slot DESC, index DESC LIMIT "limit"
					) AS ta LEFT JOIN transaction AS t ON t.signature = ta.signature AND t.partition_id = ta.partition_id;
			ELSE    
				RETURN QUERY WITH slot_getter AS (
					/* slot_filter returns the tx current slot */
					SELECT slot FROM transaction 
						WHERE signature = "current" LIMIT 1
					), 
					/* index_getter returns the current tx index */
					index_getter AS (
						SELECT index FROM transaction 
						WHERE signature = "current" LIMIT 1 
					),
					/* slot_filter includes the signature behind the current tx slot */
				slot_filter AS (
					SELECT signature, slot, index, partition_id FROM transaction_by_address WHERE address = "target" AND 
					slot <= ( SELECT slot FROM slot_getter )
				),
				/* index_filter includes the signature behind the current tx index in the current tx block */
				index_filter AS (
					SELECT signature, partition_id FROM transaction_by_address WHERE address = "target" AND 
					slot = ( SELECT slot FROM slot_getter ) AND index <= (SELECT index FROM index_getter)
				),
				/* account_signatures_getter returns the signatures filtered by the account behind the current tx slot and index */
				account_signatures_getter AS (
					SELECT slot_filter.* FROM slot_filter LEFT JOIN index_filter 
					ON slot_filter.signature = index_filter.signature AND slot_filter.partition_id = index_filter.partition_id
					ORDER BY slot DESC, index DESC LIMIT "limit" OFFSET 1 
				)   
				/* main query */  
				SELECT t.* FROM account_signatures_getter AS ta LEFT JOIN transaction AS t ON t.signature = ta.signature AND t.partition_id = ta.partition_id;
			END IF;
		END;
		$$ LANGUAGE plpgsql;;
		
		CREATE FUNCTION transactions_by_address_2(
			"target"    TEXT,
			"current"   TEXT = '',
			"limit"     INT = 10
			)
			RETURNS SETOF transaction AS
		$$
			SELECT * FROM transactions_by_address_internal("target", "current", "limit")
		$$ LANGUAGE sql STABLE;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP FUNCTION transactions_by_address_2;
	DROP FUNCTION transactions_by_address_internal;
	DROP TABLE transaction_by_address;
	`)
	return err
}
