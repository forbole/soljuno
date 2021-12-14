package postgresql

import dbtypes "github.com/forbole/soljuno/db/types"

func (db *Database) SaveInflationGovernorParam(param dbtypes.InflationGovernorParamRow) error {
	stmt := `INSERT INTO inflation_governor_param 
	(epoch, initial, terminal, taper, foundation, foundation_terminal)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		initial = excluded.initial,
		terminal = excluded.terminal,
		taper =	excluded.taper,
		foundation = excluded.foundation,
		foundation_terminal = excluded.foundation_terminal
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
		first_normal_epoch = excluded.first_normal_epoch,
		first_normal_slot = excluded.first_normal_slot,
		warmup = excluded.warmup
	WHERE epoch_schedule_param.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		param.Epoch,
		param.SlotsPerEpoch,
		param.FirstNormalEpoch,
		param.FirstNormalSlot,
		param.Warmup,
	)
	return err
}
