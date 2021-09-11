package account

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

type COption uint32

func (c COption) Bool() bool {
	return c == 1
}

type COptionPubkey struct {
	Option COption
	Value  types.Pubkey
}

type COptionUint64 struct {
	Option COption
	Value  uint64
}

type AccountState uint8

func (a AccountState) String() string {
	switch a {
	case 0:
		return "uninitialized"
	case 1:
		return "initialized"
	case 2:
		return "frozen"
	}
	return "unknown"
}

type TokenAccount struct {
	Mint           types.Pubkey
	Owner          types.Pubkey
	Amount         uint64
	Delegate       COptionPubkey
	State          AccountState
	IsNative       COptionUint64
	DelegateAmount uint64
	CloseAuthority COptionPubkey
}

func tokenParse(decoder bincode.Decoder, bz []byte) interface{} {
	switch len(bz) {
	case TokenAccountLen:
		var tokenAccount TokenAccount
		decoder.Decode(bz, &tokenAccount)
		return tokenAccount
	}
	return nil
}
