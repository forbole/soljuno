package vote

import "github.com/forbole/soljuno/solana/types"

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
