package types

type Slot uint64

type BlockConfig struct {
	Encoding                       string `json:"encoding"`
	TransactionDetails             string `json:"transactionDetails"`
	Rewards                        bool   `json:"rewards"`
	Commitment                     string `json:"commitment"`
	MaxSupportedTransactionVersion int    `json:"maxSupportedTransactionVersion"`
}

func NewDefaultBlockConfig() BlockConfig {
	return BlockConfig{
		Encoding:                       "jsonParsed",
		TransactionDetails:             "full",
		Rewards:                        true,
		Commitment:                     "finalized",
		MaxSupportedTransactionVersion: 0,
	}
}

type BlockResult struct {
	PreviousBlockhash string                             `json:"previousBlockhash"`
	Blockhash         string                             `json:"blockhash"`
	ParentSlot        Slot                               `json:"parentSlot"`
	Transactions      []EncodedTransactionWithStatusMeta `json:"transactions"`
	Rewards           []Reward                           `json:"rewards"`
	BlockTime         uint64                             `json:"blockTime"`
	BlockHeight       uint64                             `json:"blockHeight"`
}
