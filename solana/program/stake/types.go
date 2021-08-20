package stake

import "github.com/forbole/soljuno/solana/types"

type Authorized struct {
	Staker     types.Pubkey
	Withdrawer types.Pubkey
}

type Lockup struct {
	UnixTimestamp int64
	Epoch         uint64
	Custodian     types.Pubkey
}

type StakeAuthorize uint32

const (
	Staker StakeAuthorize = iota
	Withdrawer
)

type LockupArgs struct {
	LockupArgs *int64
	Epoch      *uint64
	Custodian  *types.Pubkey
}

type AuthorizeWithSeedArgs struct {
	NewAuthorizedPubkey types.Pubkey
	StakeAuthorize      StakeAuthorize
	AuthoritySeed       string
	AuthorityOwner      types.Pubkey
}

type AuthorizeCheckedWithSeedArgs struct {
	StakeAuthorize StakeAuthorize
	AuthoritySeed  string
	AuthorityOwner types.Pubkey
}

type LockupCheckedArgs struct {
	Epoch     *uint64
	Custodian *types.Pubkey
}
