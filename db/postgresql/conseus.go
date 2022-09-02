package postgresql

import (
	"fmt"
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (db *Database) GetLastBlock() (dbtypes.BlockRow, error) {
	stmt := `SELECT * FROM block ORDER BY slot DESC LIMIT 1`
	var blocks []dbtypes.BlockRow
	if err := db.Sqlx.Select(&blocks, stmt); err != nil {
		return dbtypes.BlockRow{}, err
	}

	if len(blocks) == 0 {
		return dbtypes.BlockRow{}, fmt.Errorf("cannot get block, no blocks saved")
	}

	return blocks[0], nil
}

func (db *Database) GetBlockHourAgo(now time.Time) (dbtypes.BlockRow, bool, error) {
	pastTime := now.Add(time.Hour * -1)
	return db.getBlockByTime(pastTime)
}

// -------------------------------------------------------------------------------------------------------------------

func (db *Database) SaveAverageSlotTimePerHour(slot uint64, averageTime float64) error {
	stmt := `
INSERT INTO average_slot_time_per_hour(slot, average_time) 
VALUES ($1, $2) 
ON CONFLICT (one_row_id) DO UPDATE 
    SET slot = excluded.slot, 
    average_time = excluded.average_time
WHERE average_slot_time_per_hour.slot <= excluded.slot`

	_, err := db.Sqlx.Exec(stmt, slot, averageTime)
	if err != nil {
		return fmt.Errorf("error while storing average slot time per hour: %s", err)
	}

	return nil
}
