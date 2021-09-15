package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/solana/types"
)

func stakeParse(decoder bincode.Decoder, bz []byte) interface{} {
	var stakeAccount StakeAccount
	decoder.Decode(bz, &stakeAccount)
	return stakeAccount
}

type State uint32

func (s State) String() string {
	switch s {
	case UninitializedState:
		return "uninitialized"
	case InitializedState:
		return "initialized"
	case StakeState:
		return "stake"
	case RewardsPoolState:
		return "rewardsPool"
	}
	return "unknown"
}

const (
	UninitializedState State = iota
	InitializedState
	StakeState
	RewardsPoolState
)

type StakeAccount struct {
	State State
	Meta  Meta
	Stake Stake
}

type Meta struct {
	RentExemptReserve uint64
	Authorized        stake.Authorized
	Lockup            stake.Lockup
}

type Stake struct {
	Delegation      Delegation
	CreditsObserved uint64
}

type Delegation struct {
	VoterPubkey        types.Pubkey
	Stake              uint64
	ActivationEpoch    uint64
	DeactivationEpoch  uint64
	WarmupCooldownRate float64
}
