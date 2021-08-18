package types

type Reward struct {
	Pubkey      string     `json:"pubkey"`
	Lamports    int64      `json:"lamports"`
	PostBalance uint64     `json:"postBalance"`
	RewardType  RewardType `json:"rewardType"`
	Commission  uint8      `json:"commission"`
}

type RewardType string

const (
	Fee     RewardType = "Fee"
	Rent    RewardType = "Rent"
	Staking RewardType = "Staking"
	Voting  RewardType = "Voting"
)
