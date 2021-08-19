package types

type VoteInit struct {
	NodePubkey           Pubkey
	AuthorizedVoter      Pubkey
	AuthorizedWithdrawer Pubkey
	Commission           uint8
}

type VoteAuthorize uint16

const (
	Voter VoteAuthorize = iota
	Withdrawer
)

type Vote struct {
	Slots     []uint64
	Hash      [32]byte
	Timestamp *uint64
}
