package types

type VoteAccounts struct {
	Current    []VoteAccount `json:"current"`
	Delinquent []VoteAccount `json:"delinquent"`
}

type VoteAccount struct {
	VotePubkey       string   `json:"votePubkey"`
	NodePubkey       string   `json:"nodePubkey"`
	ActivatedStake   uint64   `json:"activatedStake"`
	EpochVoteAccount bool     `json:"epochVoteAccount"`
	Commission       uint8    `json:"commission"`
	LastVote         uint64   `json:"lastVote"`
	RootSlot         uint64   `json:"rootSlot"`
	EpochCredits     [][3]int `json:"epochCredits"`
}
