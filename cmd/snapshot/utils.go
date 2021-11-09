package snapshot

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

func updateAccountBalance(ctx *Context, address string, info clienttypes.AccountInfo) error {
	bankDb, ok := ctx.Database.(db.BankDb)
	if !ok {
		return fmt.Errorf("database does not implement BankDb")
	}
	return bankDb.SaveAccountBalances(info.Context.Slot, []string{address}, []uint64{info.Value.Lamports})
}

func updateToken(ctx *Context, address string, slot uint64, token accountParser.TokenMint) error {
	tokenDb, ok := ctx.Database.(db.TokenDb)
	if !ok {
		return fmt.Errorf("database does not implement TokenDb")
	}
	return tokenDb.SaveToken(
		address,
		slot,
		token.Decimals,
		token.MintAuthority.String(),
		token.FreezeAuthority.String(),
	)
}

func updateTokenAccount(ctx *Context, address string, slot uint64, account accountParser.TokenAccount) error {
	tokenDb, ok := ctx.Database.(db.TokenDb)
	if !ok {
		return fmt.Errorf("database does not implement TokenDb")
	}
	return tokenDb.SaveTokenAccount(
		address,
		slot,
		account.Mint.String(),
		account.Owner.String(),
		"initialized",
	)
}

func updateMultisig(ctx *Context, address string, slot uint64, multisig accountParser.Multisig) error {
	tokenDb, ok := ctx.Database.(db.TokenDb)
	if !ok {
		return fmt.Errorf("database does not implement TokenDb")
	}
	return tokenDb.SaveMultisig(address, slot, multisig.Signers.Strings(), multisig.M)
}

func updateStakeAccount(ctx *Context, address string, slot uint64, account accountParser.StakeAccount) error {
	stakeDb, ok := ctx.Database.(db.StakeDb)
	if !ok {
		return fmt.Errorf("database does not implement StakeDb")
	}

	err := stakeDb.SaveStakeAccount(
		address,
		slot,
		account.Meta.Authorized.Staker.String(),
		account.Meta.Authorized.Withdrawer.String(),
		account.State.String(),
	)
	if err != nil {
		return err
	}

	err = stakeDb.SaveStakeLockup(
		address,
		slot,
		account.Meta.Lockup.Custodian.String(),
		account.Meta.Lockup.Epoch,
		account.Meta.Lockup.UnixTimestamp,
	)
	if err != nil {
		return err
	}

	delegation := account.Stake.Delegation
	return stakeDb.SaveStakeDelegation(
		address,
		slot,
		delegation.ActivationEpoch,
		delegation.DeactivationEpoch,
		delegation.Stake,
		delegation.VoterPubkey.String(),
		delegation.WarmupCooldownRate,
	)
}

func updateVoteAccount(ctx *Context, address string, slot uint64, account accountParser.VoteAccount) error {
	voteDb, ok := ctx.Database.(db.VoteDb)
	if !ok {
		return fmt.Errorf("database does not implement VoteDb")
	}
	return voteDb.SaveVoteAccount(
		address,
		slot,
		account.Node.String(),
		account.Voters[0].Pubkey.String(),
		account.Withdrawer.String(),
		account.Commission,
	)
}

func updateNonceAccount(ctx *Context, address string, slot uint64, account accountParser.NonceAccount) error {
	systemDb, ok := ctx.Database.(db.SystemDb)
	if !ok {
		return fmt.Errorf("database does not implement SystemDb")
	}
	return systemDb.SaveNonceAccount(
		address,
		slot,
		account.Authority.String(),
		account.BlockHash.String(),
		account.FeeCalculator.LamportsPerSignature,
		"initialized",
	)
}

func updateBufferAccount(ctx *Context, address string, slot uint64, account accountParser.BufferAccount) error {
	bpfLoaderDb, ok := ctx.Database.(db.BpfLoaderDb)
	if !ok {
		return fmt.Errorf("database does not implement BpfLoaderDb")
	}
	return bpfLoaderDb.SaveBufferAccount(
		address,
		slot,
		account.Authority.String(),
		"initialized",
	)
}

func updateProgramAccount(ctx *Context, address string, slot uint64, account accountParser.ProgramAccount) error {
	bpfLoaderDb, ok := ctx.Database.(db.BpfLoaderDb)
	if !ok {
		return fmt.Errorf("database does not implement BpfLoaderDb")
	}
	return bpfLoaderDb.SaveProgramAccount(
		address,
		slot,
		account.ProgramDataAccount.String(),
		"initialized",
	)
}

func updateProgramDataAccount(ctx *Context, address string, slot uint64, account accountParser.ProgramDataAccount) error {
	bpfLoaderDb, ok := ctx.Database.(db.BpfLoaderDb)
	if !ok {
		return fmt.Errorf("database does not implement BpfLoaderDb")
	}
	return bpfLoaderDb.SaveProgramDataAccount(
		address,
		slot,
		account.Slot,
		account.Authority.String(),
		"initialized",
	)
}

func updateValidatorConfig(ctx *Context, address string, slot uint64, config accountParser.ValidatorConfig) error {
	configDb, ok := ctx.Database.(db.ConfigDb)
	if !ok {
		return fmt.Errorf("database does not implement ConfigDb")
	}
	return configDb.SaveConfigAccount(
		address,
		slot,
		config.Keys[0].Pubkey.String(),
		config.Info,
	)
}
