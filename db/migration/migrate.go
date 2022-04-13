package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP FUNCTION instructions_by_address;
	CREATE FUNCTION instructions_by_address(
		addresses TEXT[],
		programs TEXT[],
		"start_slot" BIGINT = 0,
		"end_slot" BIGINT = 0
		)
		RETURNS SETOF instruction AS
	$$
	SELECT 
		instruction.tx_signature, instruction.slot, instruction.index, instruction.inner_index, instruction.program, instruction.involved_accounts, instruction.raw_data, instruction.type, instruction.value, instruction.partition_id
	FROM (
		SELECT * FROM instruction WHERE 
		(slot < "end_slot" AND slot >= "start_slot") AND
		(cardinality(programs) = 0 OR program = ANY (programs)) AND 
		involved_accounts @> addresses
		) as instruction 
	$$ LANGUAGE sql STABLE;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP FUNCTION instructions_by_address;
	CREATE FUNCTION instructions_by_address(
		addresses TEXT[],
		programs TEXT[],
		"limit" BIGINT = 100,
		"offset" BIGINT = 0)
		RETURNS SETOF instruction AS
	$$
	SELECT 
		instruction.tx_signature, instruction.slot, instruction.index, instruction.inner_index, instruction.program, instruction.involved_accounts, instruction.raw_data, instruction.type, instruction.value, instruction.partition_id
	FROM instruction
	WHERE (cardinality(programs) = 0 OR program = ANY (programs))
	  AND involved_accounts @> addresses
	ORDER BY slot DESC,
	involved_accounts LIMIT "limit" OFFSET "offset"
	$$ LANGUAGE sql STABLE;
	`)
	return err
}
