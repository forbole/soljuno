package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
)

var _ db.EpochDb = &Database{}

func (db *Database) SaveEpoch(epoch uint64) error {
	stmt := `INSERT INTO epoch (epoch) VALUES ($1)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
WHERE epoch.slot <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		epoch,
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
		foundation = exculuded.foundation,
		validator = exluced.validator
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
		circulating = exculuded.circulating,
		non_circulating = exluced.non_circulatin
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

func (db *Database) SaveInflationGovernorParam(param dbtypes.InflationGovernorParamRow) error {
	stmt := `INSERT INTO inflation_governor_param 
	(epoch, initial, terminal, taper, foundation, foundation_terminal)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		initial = excluded.initial,
		terminal = exculuded.terminal,
		taper =	excluded.taper,
		foundation = exluced.foundation,
		foundation_term = exluced.foundation_term
	WHERE inflation_governor_param.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		param.Epoch,
		param.Initial,
		param.Terminal,
		param.Taper,
		param.Foundation,
		param.FoundationTerminal,
	)
	return err
}

func (db *Database) SaveEpochScheduleParam(param dbtypes.EpochScheduleParamRow) error {
	stmt := `INSERT INTO epoch_schedule_param 
	(epoch, slots_per_epoch, first_normal_epoch, first_normal_slot, warmup)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		slots_per_epoch = excluded.slots_per_epoch,
		first_normal_epoch = exculuded.first_normal_epoch,
		first_normal_slot = exluced.first_normal_slot,
		warmup = exluced.warmup
	WHERE epoch_schedule_param.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		param.Epoch,
		param.FirstNormalEpoch,
		param.FirstNormalSlot,
		param.Warmup,
	)
	return err
}
