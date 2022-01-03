package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.EpochDb = &Database{}

func (db *Database) SaveSupplyInfo(supply dbtypes.SupplyInfoRow) error {
	stmt := `INSERT INTO supply_info 
	(epoch, total, circulating, non_circulating)
VALUES ($1, $2, $3, $4)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		total = excluded.total,
		circulating = excluded.circulating,
		non_circulating = excluded.non_circulating
	WHERE supply_info.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		supply.Epoch,
		supply.Total,
		supply.Circulating,
		supply.NonCirculating,
	)
	return err
}
