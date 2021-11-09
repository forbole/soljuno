package postgresql

import (
	"strconv"
	"time"

	"github.com/forbole/soljuno/db"
)

var _ db.StakeDb = &Database{}

// SaveStake implements the db.StakeDb
func (db *Database) SaveStakeAccount(address string, slot uint64, staker string, withdrawer string, state string) error {
	stmt := `
INSERT INTO stake_account
	(address, slot, staker, withdrawer, state)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
	SET slot = excluded.slot,
	staker = excluded.staker,
	withdrawer = excluded.withdrawer,
	state = excluded.state
WHERE stake_account.slot <= excluded.slot`

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

// SaveStakeLockup implements the db.StakeDb
func (db *Database) SaveStakeLockup(address string, slot uint64, custodian string, epoch uint64, unixTimestamp int64) error {
	stmt := `
INSERT INTO stake_lockup
    (address, slot, custodian, epoch, unix_timestamp)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    custodian = excluded.custodian,
    epoch = excluded.epoch,
    unix_timestamp = excluded.unix_timestamp
WHERE stake_lockup.slot <= excluded.slot`

	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		custodian,
		epoch,
		time.Unix(unixTimestamp, 0).UTC(),
	)
	return err
}

// SaveStakeDelegation implements the db.StakeDb
func (db *Database) SaveStakeDelegation(address string, slot uint64, activationEpoch uint64, deactivationEpoch uint64, stake uint64, voter string, rate float64) error {
	stmt := `
INSERT INTO stake_delegation
    (address, slot, activation_epoch, deactivation_epoch, stake, voter, warmup_cooldown_rate)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
    activation_epoch = excluded.activation_epoch,
    deactivation_epoch = excluded.deactivation_epoch,
    stake = excluded.stake,
    voter = excluded.voter,
    warmup_cooldown_rate = excluded.warmup_cooldown_rate
WHERE stake_delegation.slot <= excluded.slot`
	_, err := db.Sqlx.Exec(
		stmt,
		address,
		slot,
		strconv.FormatUint(activationEpoch, 10),
		strconv.FormatUint(deactivationEpoch, 10),
		stake,
		voter,
		rate,
	)
	return err
}
