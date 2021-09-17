package account_parser

import (
	"fmt"

	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

func configParse(decoder bincode.Decoder, bz []byte) interface{} {
	pos := 0
	var configAccount ConfigAccount
	var keysSize uint8
	decoder.Decode(bz[:1], &keysSize)
	pos += 1
	configAccount.Keys = make([]ConfigKey, keysSize)
	for i := 0; uint8(i) < keysSize; i++ {
		decoder.Decode(bz[pos:], &configAccount.Keys[i])
		pos += 33
	}
	fmt.Println(bz[pos:])
	decoder.Decode(bz[pos:], &configAccount.Info)
	return configAccount
}

type ConfigAccount struct {
	Keys []ConfigKey
	Info string
}

type ConfigKey struct {
	Pubkey types.Pubkey
	Signer bool
}
