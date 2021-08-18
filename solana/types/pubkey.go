package solana

import (
	"github.com/mr-tron/base58"
)

type Pubkey [32]byte

func (p Pubkey) String() string {
	return base58.Encode(p[:])
}
