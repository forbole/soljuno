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
	RewardFee     RewardType = "Fee"
	RewardRent    RewardType = "Rent"
	RewardStaking RewardType = "Staking"
	RewardVoting  RewardType = "Voting"
)
