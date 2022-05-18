package postgresql

import dbtypes "github.com/forbole/soljuno/db/types"

func (db *Database) GetMissingHeight(start uint64, end uint64) (height uint64, err error) {
	var heights []dbtypes.MissingHeightRow
	err = db.Sqlx.Select(&heights, `
	WITH start_block AS (
		SELECT slot, height FROM block WHERE slot <= $1 ORDER BY slot DESC LIMIT 1
	),
	end_block AS (
		SELECT slot, height FROM block WHERE slot >= $2 ORDER BY slot ASC LIMIT 1
	),
	missing_height AS (
		SELECT
		   generate_series AS height
		FROM
			generate_series((SELECT height FROM start_block), (SELECT height FROM end_block))
			WHERE generate_series NOT IN (
				SELECT height FROM block WHERE slot >= (SELECT slot FROM start_block) AND 
				slot <= (SELECT slot FROM end_block)
			)
	) SELECT * FROM missing_height LIMIT 1
`, start, end)
	if len(heights) == 0 {
		return 0, err
	}
	return heights[0].Height, err
}

func (db *Database) GetMissingSlotRange(height uint64) (start uint64, end uint64, err error) {
	var slots []dbtypes.MissingSlotRangeRow
	err = db.Sqlx.Select(&slots, `
	WITH start_slot AS (
		SELECT slot FROM block WHERE height <= $1 ORDER BY slot DESC LIMIT 1
	),
	end_slot AS (
		SELECT slot FROM block WHERE height >= $2 ORDER BY slot ASC LIMIT 1
	)
	SELECT 
		slot
	FROM block WHERE slot IN ( (SELECT slot FROM start_slot), (SELECT slot FROM end_slot)) ORDER BY slot ASC
	`, height, height)

	if len(slots) != 2 {
		return 0, 0, err
	}
	return slots[0].Slot + 1, slots[1].Slot - 1, err
}
