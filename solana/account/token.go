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

	case TokenLen:
		var mint Token
		decoder.Decode(bz, &mint)
		return mint

	case MultisigLen:
		var multisig Multisig
		decoder.Decode(bz, &multisig)
		return multisig
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

const TokenLen = 82

type Token struct {
	MintAuthority   COptionPubkey
	Supply          uint64
	Decimals        uint8
	IsIntialized    bool
	FreezeAuthority COptionPubkey
}

//____________________________________________________________________________

const MultisigLen = 355

type Signers [11]types.Pubkey

type Multisig struct {
	M            uint8
	N            uint8
	IsIntialized bool
	Signers      Signers
}

func (m Multisig) StringSigners() []string {
	pkStr := make([]string, m.N)
	for i := 0; i < int(m.N); i++ {
		pkStr[i] = m.Signers[i].String()
	}
	return pkStr
}
