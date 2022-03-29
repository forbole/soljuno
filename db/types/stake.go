package types

import "time"

type StakeAccountRow struct {
	Address    string `db:"address"`
	Slot       uint64 `db:"slot"`
	Staker     string `db:"staker"`
	Withdrawer string `db:"withdrawer"`
}

func NewStakeAccountRow(
	address string, slot uint64, staker string, withdrawer string,
) StakeAccountRow {
	return StakeAccountRow{
		Address:    address,
		Slot:       slot,
		Staker:     staker,
		Withdrawer: withdrawer,
	}
}

//____________________________________________________________________________

type StakeLockupRow struct {
	Address       string    `db:"address"`
	Slot          uint64    `db:"slot"`
	Custodian     string    `db:"custodian"`
	Epoch         uint64    `db:"epoch"`
	UnixTimestamp time.Time `db:"unix_timestamp"`
}

func NewStakeLockupRow(
	address string, slot uint64, custodian string, epoch uint64, unixTimestamp int64,
) StakeLockupRow {
	return StakeLockupRow{
		Address:       address,
		Slot:          slot,
		Custodian:     custodian,
		Epoch:         epoch,
		UnixTimestamp: time.Unix(unixTimestamp, 0).UTC(),
	}
}

func (lockup StakeLockupRow) Equal(other StakeLockupRow) bool {
	return lockup.Address == other.Address &&
		lockup.Slot == other.Slot &&
		lockup.Custodian == other.Custodian &&
		lockup.Epoch == other.Epoch &&
		lockup.UnixTimestamp.Equal(other.UnixTimestamp)
}

//____________________________________________________________________________

type StakeDelegationRow struct {
	Address            string  `db:"address"`
	Slot               uint64  `db:"slot"`
	ActivationEpoch    uint64  `db:"activation_epoch"`
	DeactivationEpoch  uint64  `db:"deactivation_epoch"`
	Stake              uint64  `db:"stake"`
	Voter              string  `db:"voter"`
	WarmupCooldownRate float64 `db:"warmup_cooldown_rate"`
}

func NewStakeDelegationRow(
	address string, slot uint64, activationEpoch uint64, deactivationEpoch uint64, stake uint64, voter string, rate float64,
) StakeDelegationRow {
	return StakeDelegationRow{
		Address:            address,
		Slot:               slot,
		ActivationEpoch:    activationEpoch,
		DeactivationEpoch:  deactivationEpoch,
		Stake:              stake,
		Voter:              voter,
		WarmupCooldownRate: rate,
	}
}
