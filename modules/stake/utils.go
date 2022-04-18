package stake

import (
	"encoding/base64"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
)

// updateStakeAccount properly stores the statement of stake account inside the database
func updateStakeAccount(address string, currentSlot uint64, db db.StakeDb, client ClientProxy) error {
	if db.CheckStakeAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteStakeAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	stakeAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.StakeAccount)
	if !ok {
		return db.DeleteStakeAccount(address)
	}

	err = db.SaveStakeAccount(
		dbtypes.NewStakeAccountRow(
			address,
			info.Context.Slot,
			stakeAccount.Meta.Authorized.Staker.String(),
			stakeAccount.Meta.Authorized.Withdrawer.String(),
		),
	)
	if err != nil {
		return err
	}

	err = db.SaveStakeLockup(
		dbtypes.NewStakeLockupRow(
			address,
			info.Context.Slot,
			stakeAccount.Meta.Lockup.Custodian.String(),
			stakeAccount.Meta.Lockup.Epoch,
			stakeAccount.Meta.Lockup.UnixTimestamp,
		),
	)
	if err != nil {
		return err
	}

	if stakeAccount.State.String() != "stake" {
		return db.DeleteStakeDelegation(address)
	}

	delegation := stakeAccount.Stake.Delegation
	return db.SaveStakeDelegation(
		dbtypes.NewStakeDelegationRow(
			address,
			info.Context.Slot,
			delegation.ActivationEpoch,
			delegation.DeactivationEpoch,
			delegation.Stake,
			delegation.VoterPubkey.String(),
			delegation.WarmupCooldownRate,
		),
	)
}
