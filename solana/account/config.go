package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/types"
)

type ConfigType uint8

const (
	StakeConfigType ConfigType = iota
)

func configParse(decoder bincode.Decoder, bz []byte) interface{} {
	var typ ConfigType
	decoder.Decode(bz[0:1], &typ)
	switch typ {
	case StakeConfigType:
		return parseStakeConfig(decoder, bz[1:])
	default:
		return parseValidatorConfig(decoder, bz[0:])
	}
}

func parseStakeConfig(decoder bincode.Decoder, bz []byte) interface{} {
	var config StakeConfig
	decoder.Decode(bz, &config)
	return config
}

func parseValidatorConfig(decoder bincode.Decoder, bz []byte) interface{} {
	var config ValidatorConfig
	var keysSize uint8
	decoder.Decode(bz[:1], &keysSize)
	config.Keys = make([]ConfigKey, keysSize)
	pos := 1
	for i := 0; uint8(i) < keysSize; i++ {
		decoder.Decode(bz[pos:pos+33], &config.Keys[i])
		pos += 33
	}
	decoder.Decode(bz[pos:], &config.Info)
	return config
}

type ValidatorConfig struct {
	Keys []ConfigKey
	Info string
}

type ConfigKey struct {
	Pubkey types.Pubkey
	Signer bool
}

type StakeConfig struct {
	WarmupCooldownRate float64
	SlashPenalty       uint8
}
