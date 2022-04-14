package snapshot

import (
	"encoding/json"

	"github.com/forbole/soljuno/apis/keybase"
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
)

func saveToken(db db.TokenDb, address string, slot uint64, token parser.Token) error {
	err := db.SaveToken(
		dbtypes.NewTokenRow(
			address,
			slot,
			token.Decimals,
			token.MintAuthority.String(),
			token.FreezeAuthority.String(),
		),
	)
	if err != nil {
		return err
	}
	return db.SaveTokenSupply(dbtypes.NewTokenSupplyRow(address, slot, token.Supply))
}

func saveTokenAccount(db db.TokenDb, address string, slot uint64, account parser.TokenAccount) error {
	return db.SaveTokenAccount(
		dbtypes.NewTokenAccountRow(
			address,
			slot,
			account.Mint.String(),
			account.Owner.String(),
		),
	)
}

func saveTokenBalance(db db.BankDb, address string, slot uint64, account parser.TokenAccount) error {
	return db.SaveAccountTokenBalances(
		slot,
		[]string{address},
		[]uint64{account.Amount},
	)
}

func saveMultisig(db db.TokenDb, address string, slot uint64, multisig parser.Multisig) error {
	return db.SaveMultisig(
		dbtypes.NewMultisigRow(
			address, slot, multisig.StringSigners(), multisig.M,
		),
	)
}

func updateStakeAccount(db db.StakeDb, address string, slot uint64, account parser.StakeAccount) error {
	err := db.SaveStakeAccount(
		dbtypes.NewStakeAccountRow(
			address,
			slot,
			account.Meta.Authorized.Staker.String(),
			account.Meta.Authorized.Withdrawer.String(),
		),
	)
	if err != nil {
		return err
	}

	err = db.SaveStakeLockup(
		dbtypes.NewStakeLockupRow(
			address,
			slot,
			account.Meta.Lockup.Custodian.String(),
			account.Meta.Lockup.Epoch,
			account.Meta.Lockup.UnixTimestamp,
		),
	)
	if err != nil {
		return err
	}

	delegation := account.Stake.Delegation
	return db.SaveStakeDelegation(
		dbtypes.NewStakeDelegationRow(
			address,
			slot,
			delegation.ActivationEpoch,
			delegation.DeactivationEpoch,
			delegation.Stake,
			delegation.VoterPubkey.String(),
			delegation.WarmupCooldownRate,
		),
	)
}

func saveVoteAccount(db db.VoteDb, address string, slot uint64, account parser.VoteAccount) error {
	return db.SaveValidator(
		dbtypes.NewVoteAccountRow(
			address,
			slot,
			account.Node.String(),
			account.Voters[0].Pubkey.String(),
			account.Withdrawer.String(),
			account.Commission,
		),
	)
}

func updateNonceAccount(db db.SystemDb, address string, slot uint64, account parser.NonceAccount) error {
	return db.SaveNonceAccount(
		dbtypes.NewNonceAccountRow(
			address,
			slot,
			account.Authority.String(),
			account.BlockHash.String(),
			account.FeeCalculator.LamportsPerSignature,
		),
	)
}

func updateBufferAccount(db db.BpfLoaderDb, address string, slot uint64, account parser.BufferAccount) error {
	return db.SaveBufferAccount(
		dbtypes.NewBufferAccountRow(
			address,
			slot,
			account.Authority.String(),
		),
	)
}

func updateProgramAccount(db db.BpfLoaderDb, address string, slot uint64, account parser.ProgramAccount) error {
	return db.SaveProgramAccount(
		dbtypes.NewProgramAccountRow(
			address,
			slot,
			account.ProgramDataAccount.String(),
		),
	)
}

func updateProgramDataAccount(db db.BpfLoaderDb, address string, slot uint64, account parser.ProgramDataAccount) error {
	return db.SaveProgramDataAccount(
		dbtypes.NewProgramDataAccountRow(
			address,
			slot,
			account.Slot,
			account.Authority.String(),
		),
	)
}

func updateValidatorConfig(db db.ConfigDb, address string, slot uint64, config parser.ValidatorConfig) error {
	var parsedConfig dbtypes.ParsedValidatorConfig
	err := json.Unmarshal([]byte(config.Info), &parsedConfig)
	if err != nil {
		return err
	}

	kbClient := keybase.NewClient()
	avatarUrl, err := kbClient.GetAvatarURL(parsedConfig.KeybaseUsername)
	if err != nil {
		avatarUrl = ""
	}

	row := dbtypes.NewValidatorConfigRow(
		address, slot, config.Keys[1].Pubkey.String(), parsedConfig, avatarUrl,
	)

	return db.SaveValidatorConfig(row)
}
