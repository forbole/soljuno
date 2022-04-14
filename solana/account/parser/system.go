package parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

func systemParse(decoder bincode.Decoder, bz []byte) interface{} {
	if len(bz) != 0 {
		var nonce NonceAccount
		decoder.Decode(bz, &nonce)
		return nonce
	}
	return nil
}

type NonceState uint32

const (
	NonceUninitialized NonceState = iota
	NonceInitialized
)

func (n NonceState) String() string {
	if n == NonceUninitialized {
		return "uninitialized"
	}
	return "initialized"
}

type NonceAccount struct {
	Current       uint32
	State         NonceState
	Authority     types.Pubkey
	BlockHash     types.Hash
	FeeCalculator types.FeeCalculator
}
