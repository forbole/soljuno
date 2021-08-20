package types

import "github.com/mr-tron/base58"

type Hash [32]byte

func (h Hash) String() string {
	return base58.Encode(h[:])
}
