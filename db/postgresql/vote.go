package postgresql

import "github.com/forbole/soljuno/db"

var _ db.VoteDb = &Database{}

// SaveVoteAccount implements the db.VoteDb
func (db *Database) SaveVoteAccount(address string, slot uint64, node string, voter string, withdrawer string, commission uint8) error {
	stmt := `
INSERT INTO vote_account
    (address, slot, node, voter, withdrawer, commission)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    node = excluded.node,
	voter = excluded.voter,
	withdrawer = excluded.withdrawer,
	commission = excluded.commission
WHERE vote_account.slot <= excluded.slot`
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
ON CONFLICT (address, slot) DO NOTHING`
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

// PruneValidatorStatus implements the db.VoteDb
func (db *Database) PruneValidatorStatus(slot uint64) error {
	_, err := db.Sqlx.Exec(`DELETE FROM validator_status WHERE slot <= $1`, slot)
	return err
}
