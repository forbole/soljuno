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

// SaveValidatorStatus implements the db.VoteDb
func (db *Database) SaveValidatorStatus(address string, slot uint64, activatedStake uint64, lastVote uint64, rootSlot uint64, active bool) error {
	stmt := `
INSERT INTO validator_status
	(address, slot, activated_stake, last_vote, root_slot, active)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (address) DO UPDATE 
	SET slot = EXCLUDED.slot,
	    activated_stake = EXCLUDED.activated_stake,
		last_vote = EXCLUDED.last_vote,
		root_slot = EXCLUDED.root_slot,
		active = EXCLUDED.active
	WHERE validator_status.slot <= EXCLUDED.slot
		`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		activatedStake,
		lastVote,
		rootSlot,
		active,
	)
	return err
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
	insertStmt := `INSERT INTO validator_skip_rate (address, epoch, skip_rate) VALUES`
	paramsStmt := ""
	conflictStmt := `
	ON CONFLICT (address) DO UPDATE
		SET epoch = excluded.epoch,
			skip_rate = excluded.skip_rate
	WHERE validator_skip_rate.epoch <= excluded.epoch
	`
	paramsNumber := 3
	var params []interface{}
	for i, row := range skipRates {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(params, row.Address, row.Epoch, row.SkipRate)
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}
