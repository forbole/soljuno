package stake

import (
	"encoding/base64"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
)

// updateStakeAccount properly stores the statement of stake account inside the database
func updateStakeAccount(address string, currentSlot uint64, db db.StakeDb, client client.Proxy) error {
	if !db.CheckStakeAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveStakeAccount(
			address,
			info.Context.Slot,
			"",
			"",
			"closed",
		)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	stakeAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.StakeAccount)
	if !ok {
		return db.SaveStakeAccount(
			address,
			info.Context.Slot,
			"",
			"",
			"closed",
		)
	}

	err = db.SaveStakeAccount(
		address,
		info.Context.Slot,
		stakeAccount.Meta.Authorized.Staker.String(),
		stakeAccount.Meta.Authorized.Withdrawer.String(),
		stakeAccount.State.String(),
	)
	if err != nil {
		return err
	}

	err = db.SaveStakeLockup(
		address,
		info.Context.Slot,
		stakeAccount.Meta.Lockup.Custodian.String(),
		stakeAccount.Meta.Lockup.Epoch,
		stakeAccount.Meta.Lockup.UnixTimestamp,
	)
	if err != nil {
		return err
	}

	if stakeAccount.State.String() == "stake" {
		delegation := stakeAccount.Stake.Delegation
		err := db.SaveStakeDelegation(
			address,
			info.Context.Slot,
			delegation.ActivationEpoch,
			delegation.DeactivationEpoch,
			delegation.Stake,
			delegation.VoterPubkey.String(),
			delegation.WarmupCooldownRate,
		)
		if err != nil {
			return err
		}
	}

	return db.SaveStakeDelegation(
		address,
		info.Context.Slot,
		0,
		0,
		0,
		"",
		0,
	)
}
