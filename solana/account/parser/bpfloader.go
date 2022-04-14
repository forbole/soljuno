package parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

func bpfLoaderParse(decoder bincode.Decoder, bz []byte) interface{} {
	var state LoaderState
	decoder.Decode(bz[:4], &state)
	switch state {
	case BufferLoaderState:
		return bufferAccountParse(decoder, bz[4:])
	case ProgramLoaderState:
		return programAccountParse(decoder, bz[4:])
	case ProgramDataLoaderState:
		return programDataAccountParse(decoder, bz[4:])
	}
	return nil
}

func bufferAccountParse(decoder bincode.Decoder, bz []byte) interface{} {
	var account BufferAccount
	decoder.Decode(bz, &account)
	return account
}

func programAccountParse(decoder bincode.Decoder, bz []byte) interface{} {
	var account ProgramAccount
	decoder.Decode(bz, &account)
	return account
}

func programDataAccountParse(decoder bincode.Decoder, bz []byte) interface{} {
	var account ProgramDataAccount
	decoder.Decode(bz, &account)
	return account
}

type LoaderState uint32

const (
	UninitializedLoaderState LoaderState = iota
	BufferLoaderState
	ProgramLoaderState
	ProgramDataLoaderState
)

type BufferAccount struct {
	Authority types.Pubkey
}

type ProgramAccount struct {
	ProgramDataAccount types.Pubkey
}

type ProgramDataAccount struct {
	Slot      uint64
	Authority types.Pubkey
}
