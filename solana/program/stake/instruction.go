package stake

import "github.com/forbole/soljuno/solana/types"

type InstructionID uint32

const (
	// Initialize a stake with lockup and authorization information
	Initialize InstructionID = iota

	// Authorize a key to manage stake or withdrawal
	Authorize

	// Delegate a stake to a particular vote account
	DelegateStake

	// Split u64 tokens and stake off a stake account into another stake account
	Split

	// Withdraw unstaked lamports from the stake account
	Withdraw

	// Deactivates the stake in the account
	Deactivate

	// Set stake lockup
	SetLockup

	// Authorize a key to manage stake or withdrawal with a derived key
	Merge

	// Initialize a stake with authorization information
	AuthorizeWithSeed

	// Initialize a stake with authorization information
	InitializeChecked

	// Authorize a key to manage stake or withdrawal
	AuthorizeChecked

	// Authorize a key to manage stake or withdrawal with a derived key
	AuthorizeCheckedWithSeed

	// Set stake lockup
	SetLockupChecked
)

type InitializeInstruction struct {
	Authorized Authorized
	Lockup     Lockup
}

type AuthorizeInstruction struct {
	Pubkey         types.Pubkey
	StakeAuthorize StakeAuthorize
}

type SplitInstruction struct {
	Amount uint64
}

type WithdrawInstruction struct {
	Amount uint64
}

type SetLockupInstruction struct {
	LockupArgs LockupArgs
}

type AuthorizeWithSeedInstruction struct {
	AuthorizeWithSeedArgs AuthorizeWithSeedArgs
}

type AuthorizeCheckedInstruction struct {
	StakeAuthorize StakeAuthorize
}

type AuthorizeCheckedWithSeedInstruction struct {
	AuthorizeCheckedWithSeedArgs AuthorizeCheckedWithSeedArgs
}

type SetLockupCheckedInstruction struct {
	LockupCheckedArgs LockupCheckedArgs
}

type Authorized struct {
	Staker     types.Pubkey
	Withdrawer types.Pubkey
}

type Lockup struct {
	UnixTimestamp int64
	Epoch         uint64
	Custodian     types.Pubkey
}

type StakeAuthorize uint32

const (
	Staker StakeAuthorize = iota
	Withdrawer
)

type LockupArgs struct {
	LockupArgs *int64
	Epoch      *uint64
	Custodian  *types.Pubkey
}

type AuthorizeWithSeedArgs struct {
	NewAuthorizedPubkey types.Pubkey
	StakeAuthorize      StakeAuthorize
	AuthoritySeed       string
	AuthorityOwner      types.Pubkey
}

type AuthorizeCheckedWithSeedArgs struct {
	StakeAuthorize StakeAuthorize
	AuthoritySeed  string
	AuthorityOwner types.Pubkey
}

type LockupCheckedArgs struct {
	Epoch     *uint64
	Custodian *types.Pubkey
}
