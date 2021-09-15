package postgresql

import "github.com/forbole/soljuno/db"

var _ db.StakeDb = &Database{}

func (db *Database) SaveStake(address string, slot uint64, staker string, withdrawer string, state string) error {
	stmt := `
INSERT INTO stake
	(address, slot, staker, withdrawer, state)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	staker = excluded.staker,
	withdrawer = excluded.withdrawer,
	state = excluded.state
WHERE nonce.slot <= excluded.slot`

	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		staker,
		withdrawer,
		state,
	)
	return err
}

func (db *Database) SaveStakeLockup(address string, slot uint64, custodian string, epoch uint64, unixTimestamp uint64) error {
	stmt := `
INSERT INTO stake_lockup
    (address, slot, custodian, epoch, unixTimestamp)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    custodian = excluded.custodian,
    epoch = excluded.epoch,
    unixTimestamp = excluded.unixTimestamp
WHERE nonce.slot <= excluded.slot`

	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		custodian,
		epoch,
		unixTimestamp,
	)
	return err
}

func (db *Database) SaveStakeDelegation(address string, slot uint64, activationEpoch uint64, deactivationEpoch uint64, stake uint64, voter string, rate float64) error {
	stmt := `
INSERT INTO stake_delegation
    (address, slot, activationEpoch, deactivationEpoch, stake, voter, warmup_cooldown_rate)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    activationEpoch = excluded.activationEpoch,
    blockhash = excluded.blockhash,
    deactivationEpoch = excluded.deactivationEpoch,
    stake = excluded.stake,
    voter = excluded.voter,
    warmup_cooldown_rate = excluded.warmup_cooldown_rate
WHERE nonce.slot <= excluded.slot`

	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		activationEpoch,
		deactivationEpoch,
		stake,
		voter,
		rate,
	)
	return err
}
