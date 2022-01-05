package postgresql

import "github.com/forbole/soljuno/db"

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
