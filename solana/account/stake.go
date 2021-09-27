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

type StakeState uint32

func (s StakeState) String() string {
	switch s {
	case UninitializedStakeState:
		return "uninitialized"
	case InitializedStakeState:
		return "initialized"
	case StakedStakeState:
		return "stake"
	case RewardsPoolStakeState:
		return "rewardsPool"
	}
	return "unknown"
}

const (
	UninitializedStakeState StakeState = iota
	InitializedStakeState
	StakedStakeState
	RewardsPoolStakeState
)

type StakeAccount struct {
	State StakeState
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
