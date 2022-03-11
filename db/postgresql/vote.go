package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	solanatypes "github.com/forbole/soljuno/solana/types"
)

var _ db.VoteDb = &Database{}

// SaveValidator implements the db.VoteDb
func (db *Database) SaveValidator(address string, slot uint64, node string, voter string, withdrawer string, commission uint8) error {
	stmt := `
INSERT INTO validator
    (address, slot, node, voter, withdrawer, commission)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    node = excluded.node,
	voter = excluded.voter,
	withdrawer = excluded.withdrawer,
	commission = excluded.commission
WHERE validator.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		node,
		voter,
		withdrawer,
		commission,
	)
	return err
}

// SaveValidatorStatuses implements the db.VoteDb
func (db *Database) SaveValidatorStatuses(statuses []dbtypes.ValidatorStatusRow) error {
	// clean the current state
	if _, err := db.Sqlx.Exec("TRUNCATE validator_status"); err != nil {
		return err
	}

	stmt := `
INSERT INTO validator_status
	(address, slot, activated_stake, last_vote, root_slot, active) VALUES`
	var params []interface{}
	paramsNumber := 6
	params = make([]interface{}, 0, paramsNumber*len(statuses))
	for _, status := range statuses {
		params = append(params, status.Address, status.Slot, status.ActivatedStake, status.LastVote, status.Active)
	}
	return db.InsertBatch(
		stmt,
		"",
		params,
		paramsNumber,
	)
}

// GetEpochProducedBlocks implements the db.VoteDb
func (db *Database) GetEpochProducedBlocks(epoch uint64) ([]uint64, error) {
	start := epoch * solanatypes.SlotsInEpoch
	end := (epoch+1)*solanatypes.SlotsInEpoch - 1
	var blocks []dbtypes.BlockRow
	err := db.Sqlx.Select(&blocks, `SELECT slot FROM block WHERE slot >= $1 AND slot <= $2`, start, end)
	if err != nil {
		return nil, err
	}
	res := make([]uint64, len(blocks))
	for i, block := range blocks {
		res[i] = block.Slot
	}
	return res, nil
}

// SaveValidatorSkipRates implements the db.VoteDb
func (db *Database) SaveValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error {
	insertStmt := `INSERT INTO validator_skip_rate (address, epoch, skip_rate, total, skip) VALUES`
	conflictStmt := `
	ON CONFLICT (address) DO UPDATE
		SET epoch = excluded.epoch,
			skip_rate = excluded.skip_rate,
			total = excluded.total,
			skip = excluded.skip
	WHERE validator_skip_rate.epoch <= excluded.epoch
	`
	var params []interface{}
	paramsNumber := 5
	params = make([]interface{}, 0, paramsNumber*len(skipRates))
	for _, row := range skipRates {
		params = append(params, row.Address, row.Epoch, row.SkipRate, row.Total, row.Skip)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// SaveValidatorSkipRates implements the db.VoteDb
func (db *Database) SaveHistoryValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error {
	insertStmt := `INSERT INTO validator_skip_rate_history (address, epoch, skip_rate, total, skip) VALUES`
	conflictStmt := `
	ON CONFLICT (address, epoch) DO NOTHING
	`
	var params []interface{}
	paramsNumber := 5
	params = make([]interface{}, 0, paramsNumber*len(skipRates))
	for _, row := range skipRates {
		params = append(params, row.Address, row.Epoch, row.SkipRate, row.Total, row.Skip)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}
