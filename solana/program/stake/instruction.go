package stake

import "github.com/cosmos/cosmos-sdk/crypto/types"

type StakeInstructionID uint32

const (
	// Initialize a stake with lockup and authorization information
	Initialize StakeInstructionID = iota

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
	Pubkey         types.PubKey
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
