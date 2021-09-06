package vote

import "github.com/forbole/soljuno/solana/types"

type InstructionID uint32

const (
	// Initialize a vote account
	InitializeAccount InstructionID = iota

	// Authorize a key to send votes or issue a withdrawal
	Authorize

	// A Vote instruction with recent votes
	Vote

	// Withdraw some amount of funds
	Withdraw

	// Update the vote account's validator identity (node_pubkey)
	UpdateValidatorIdentity

	// Update the commission for the vote account
	UpdateCommission

	// A Vote instruction with recent votes
	VoteSwitch

	// Authorize a key to send votes or issue a withdrawal
	// This instruction behaves like `Authorize` with the additional requirement that the new vote
	// or withdraw authority must also be a signer.
	AuthorizeChecked
)

//____________________________________________________________________________

// Instruction struct

type InitializeAccountInstruction struct {
	VoteInit VoteInit
}

type AuthorizeInstruction struct {
	Pubkey        types.Pubkey
	VoteAuthorize VoteAuthorize
}

type VoteInstruction struct {
	Vote VoteData
}

type WithdrawInstruction struct {
	Amount uint64
}

type UpdateCommissionInstruction struct {
	Commission uint8
}

type VoteSwitchInstruction struct {
	Vote VoteData
	Hash types.Hash
}

type AuthorizeCheckedInstruction struct {
	VoteAuthorize VoteAuthorize
}

//____________________________________________________________________________

// The instance used by instructions

type VoteInit struct {
	NodePubkey           types.Pubkey
	AuthorizedVoter      types.Pubkey
	AuthorizedWithdrawer types.Pubkey
	Commission           uint8
}

type VoteAuthorize uint16

const (
	Voter VoteAuthorize = iota
	Withdrawer
)

type VoteData struct {
	Slots     []uint64   `json:"slots"`
	Hash      types.Hash `json:"hash"`
	Timestamp *uint64    `json:"timestamp,omitempty"`
}
