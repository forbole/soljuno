package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

func voteParse(decoder bincode.Decoder, bz []byte) interface{} {
	if len(bz) != 0 {
		var voteAccount VoteAccount
		decoder.Decode(bz, &voteAccount)
		return voteAccount
	}
	return nil
}

type VoteAccount struct {
	Node           types.Pubkey
	Withdrawer     types.Pubkey
	Commission     uint8
	Votes          []Lockout
	RootSlot       *uint64
	Voter          []Voter
	PriorVoters    [32]PriorVoter
	CircBufCfg     CircBufCfg
	EpochCredits   []EpochCredit
	BlockTimestamp BlockTimestamp
}

type Lockout struct {
	Slot              uint64
	ConfirmationCount uint32
}

type Voter struct {
	Epoch  uint64
	Pubkey types.Pubkey
}

type CircBufCfg struct {
	Next    uint64
	IsEmpty bool
}

type PriorVoter struct {
	Pubkey     types.Pubkey
	StartEpoch uint64
	EndEpoch   uint64
}

type EpochCredit struct {
	Epoch      uint64
	Credit     uint64
	PrevCredit uint64
}

type BlockTimestamp struct {
	Slot          uint64
	UnixTimestamp int64
}
