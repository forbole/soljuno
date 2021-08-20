package system

import "github.com/forbole/soljuno/solana/types"

type SystemInstructionID uint32

const (
	// Create a new account
	CreateAccount SystemInstructionID = iota

	// Assign account to a program
	Assign

	// Transfer lamports
	Transfer

	// Create a new account at an address derived from a base pubkey and a seed
	CreateAccountWithSeed

	// Consumes a stored nonce, replacing it with a successor
	AdvanceNonceAccount

	// Withdraw funds from a nonce account
	WithdrawNonceAccount

	// Drive state of Uninitialized nonce account to Initialized, setting the nonce value
	InitializeNonceAccount

	// Change the entity authorized to execute nonce instructions on the account
	AuthorizeNonceAccount

	// Allocate space in a (possibly new) account without funding
	Allocate

	// Allocate space for and assign an account at an address
	AllocateWithSeed

	// Assign account to a program based on a seed
	AssignWithSeed

	// Transfer lamports from a derived address
	TransferWithSeed
)

type CreateAccountInstruction struct {
	Lamports uint64
	Space    uint64
	Owner    types.Pubkey
}

type AssignInstruction struct {
	Owner types.Pubkey
}

type TransferInstruction struct {
	Lamports uint64
}

type CreateAccountWithSeedInstruction struct {
	Base     types.Pubkey
	Seed     string
	Lamports uint64
	Space    uint64
	Owner    types.Pubkey
}

type WithdrawNonceAccountInstruction struct {
	Amount uint64
}

type InitializeNonceAccountInstruction struct {
	Account types.Pubkey
}

type AuthorizeNonceAccountInstruction struct {
	Account types.Pubkey
}

type AllocateInstruction struct {
	Space uint64
}

type AllocateWithSeedInstruction struct {
	Base  types.Pubkey
	Seed  string
	Space uint64
	Owner types.Pubkey
}

type AssignWithSeedInstruction struct {
	Base  types.Pubkey
	Seed  string
	Owner types.Pubkey
}

type TransferWithSeedInstruction struct {
	Lamports  uint64
	FromSeed  string
	FromOwner types.Pubkey
}
