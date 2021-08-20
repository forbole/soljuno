package token

import "github.com/forbole/soljuno/solana/types"

type TokenInstructionID uint32

const (
	// Initializes a new mint and optionally deposits all the newly minted
	InitializeMint TokenInstructionID = iota

	// Initializes a new account to hold tokens
	InitializeAccount

	// Initializes a multisignature account with N provided signers
	InitializeMultisig

	// Transfers tokens from one account to another either directly or via a delegate
	Transfer

	// Approves a delegate
	Approve

	// Revokes the delegate's authority
	Revoke

	// Sets a new authority of a mint or account
	SetAuthority

	// Mints new tokens to an account
	MintTo

	// Burns tokens by removing them from an account
	Burn

	// Close an account by transferring all its SOL to the destination account
	CloseAccount

	// Freeze an Initialized account using the Mint's freeze_authority (if it set)
	FreezeAccount

	// Thaw a Frozen account using the Mint's freeze_authority (if set)
	ThawAccount

	// Transfers tokens from one account to another either directly or via a delegate
	TransferChecked

	// Approves a delegate
	// This instruction differs from Approve in that the token mint and decimals value is checked by the caller
	ApproveChecked

	// Mints new tokens to an account
	// This instruction differs from MintTo in that the decimals value is checked by the caller
	MintToChecked

	// Burns tokens by removing them from an account
	BurnChecked

	// Like InitializeAccount, but the owner pubkey is passed via instruction data rather than the accounts
	InitializeAccount2

	SyncNative

	// Like InitializeAccount2, but does not require the Rent sysvar to be provided
	InitializeAccount3

	// Like InitializeMultisig, but does not require the Rent sysvar to be provided
	InitializeMultisig2

	// Like InitializeMint, but does not require the Rent sysvar to be provided
	InitializeMint2
)

type InitializeMintInstruction struct {
	Decimals        uint8
	MintAuthority   types.Pubkey
	FreezeAuthority *types.Pubkey
}

type InitializeMultisigInstruction struct {
	Amount uint8
}

type TransferInstruction struct {
	Amount uint64
}

type ApproveInstruction struct {
	Amount uint64
}

type SetAuthorityInstruction struct {
	AuthorityType AuthorityType
}

type MintToInstruction struct {
	Amount uint64
}

type BurnInstruction struct {
	Amount uint64
}

type TransferCheckedInstruction struct {
	Amount   uint64
	Decimals uint8
}

type ApproveCheckedInstruction struct {
	Amount   uint64
	Decimals uint8
}

type MintToCheckedInstruction struct {
	Amount   uint64
	Decimals uint8
}

type BurnCheckedInstruction struct {
	Amount   uint64
	Decimals uint8
}

type InitializeAccountInstruction struct {
	Owner types.Pubkey
}
