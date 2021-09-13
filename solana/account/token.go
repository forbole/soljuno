package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

func tokenParse(decoder bincode.Decoder, bz []byte) interface{} {
	switch len(bz) {
	case TokenAccountLen:
		var account TokenAccount
		decoder.Decode(bz, &account)
		return account

	case TokenMintLen:
		var mint TokenMint
		decoder.Decode(bz, &mint)
		return mint
	}
	return nil
}

//____________________________________________________________________________

type COption uint32

func (c COption) Bool() bool {
	return c == 1
}

type COptionPubkey struct {
	Option COption
	Value  types.Pubkey
}

func (c COptionPubkey) String() string {
	if !c.Option.Bool() {
		return ""
	}
	return c.Value.String()
}

type COptionUint64 struct {
	Option COption
	Value  uint64
}

//____________________________________________________________________________

const TokenAccountLen = 165

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

//____________________________________________________________________________

const TokenMintLen = 82

type TokenMint struct {
	MintAuthority   COptionPubkey
	Supply          uint64
	Decimals        uint8
	IsIntialized    bool
	FreezeAuthority COptionPubkey
}
