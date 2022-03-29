package snapshot

import (
	"encoding/json"

	"github.com/forbole/soljuno/apis/keybase"
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	accountParser "github.com/forbole/soljuno/solana/account"
)

func updateToken(ctx *Context, address string, slot uint64, token accountParser.Token) error {
	tokenDb := ctx.Database.(db.TokenDb)
	err := tokenDb.SaveToken(
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

	return tokenDb.SaveTokenSupply(dbtypes.NewTokenSupplyRow(address, slot, token.Supply))
}

func updateTokenAccount(ctx *Context, address string, slot uint64, account accountParser.TokenAccount) error {
	tokenDb := ctx.Database.(db.TokenDb)
	err := tokenDb.SaveTokenAccount(
		dbtypes.NewTokenAccountRow(
			address,
			slot,
			account.Mint.String(),
			account.Owner.String(),
		),
	)
	if err != nil {
		return err
	}

	bankDb := ctx.Database.(db.BankDb)
	return bankDb.SaveAccountTokenBalances(
		slot,
		[]string{address},
		[]uint64{account.Amount},
	)
}

func updateMultisig(ctx *Context, address string, slot uint64, multisig accountParser.Multisig) error {
	tokenDb := ctx.Database.(db.TokenDb)
	return tokenDb.SaveMultisig(
		dbtypes.NewMultisigRow(
			address, slot, multisig.StringSigners(), multisig.M,
		),
	)
}

func updateStakeAccount(ctx *Context, address string, slot uint64, account accountParser.StakeAccount) error {
	stakeDb := ctx.Database.(db.StakeDb)
	err := stakeDb.SaveStakeAccount(
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

	err = stakeDb.SaveStakeLockup(
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
	return stakeDb.SaveStakeDelegation(
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

func updateVoteAccount(ctx *Context, address string, slot uint64, account accountParser.VoteAccount) error {
	voteDb := ctx.Database.(db.VoteDb)
	return voteDb.SaveValidator(
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

func updateNonceAccount(ctx *Context, address string, slot uint64, account accountParser.NonceAccount) error {
	systemDb := ctx.Database.(db.SystemDb)
	return systemDb.SaveNonceAccount(
		dbtypes.NewNonceAccountRow(
			address,
			slot,
			account.Authority.String(),
			account.BlockHash.String(),
			account.FeeCalculator.LamportsPerSignature,
		),
	)
}

func updateBufferAccount(ctx *Context, address string, slot uint64, account accountParser.BufferAccount) error {
	bpfLoaderDb := ctx.Database.(db.BpfLoaderDb)
	return bpfLoaderDb.SaveBufferAccount(
		dbtypes.NewBufferAccountRow(
			address,
			slot,
			account.Authority.String(),
		),
	)
}

func updateProgramAccount(ctx *Context, address string, slot uint64, account accountParser.ProgramAccount) error {
	bpfLoaderDb := ctx.Database.(db.BpfLoaderDb)
	return bpfLoaderDb.SaveProgramAccount(
		dbtypes.NewProgramAccountRow(
			address,
			slot,
			account.ProgramDataAccount.String(),
		),
	)
}

func updateProgramDataAccount(ctx *Context, address string, slot uint64, account accountParser.ProgramDataAccount) error {
	bpfLoaderDb := ctx.Database.(db.BpfLoaderDb)
	return bpfLoaderDb.SaveProgramDataAccount(
		dbtypes.NewProgramDataAccountRow(
			address,
			slot,
			account.Slot,
			account.Authority.String(),
		),
	)
}

func updateValidatorConfig(ctx *Context, address string, slot uint64, config accountParser.ValidatorConfig) error {
	configDb := ctx.Database.(db.ConfigDb)
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

	return configDb.SaveValidatorConfig(row)
}
