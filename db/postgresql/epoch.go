package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.EpochDb = &Database{}

func (db *Database) SaveEpoch(info dbtypes.EpochInfoRow) error {
	stmt := `INSERT INTO epoch_info (epoch, transaction_count) VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch
WHERE epoch_info.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		info.Epoch,
		info.TransactionCount,
	)
	return err
}

func (db *Database) SaveInflationRate(inflation dbtypes.InflationRateRow) error {
	stmt := `INSERT INTO inflation_rate 
	(epoch, total, foundation, validator)
VALUES ($1, $2, $3, $4)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		total = excluded.total,
		foundation = excluded.foundation,
		validator = excluded.validator
	WHERE inflation_rate.epoch <= excluded.epoch
`
	_, err := db.Sqlx.Exec(
		stmt,
		inflation.Epoch,
		inflation.Total,
		inflation.Foundation,
		inflation.Validator,
	)
	return err
}

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
