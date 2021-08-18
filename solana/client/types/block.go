package types

type Slot uint64

type Block struct {
	PreviousBlockhash string                             `json:"previousBlockhash"`
	Blockhash         string                             `json:"blockhash"`
	ParentSlot        Slot                               `json:"parentSlot"`
	Transactions      []EncodedTransactionWithStatusMeta `json:"transactions"`
	Rewards           []Reward                           `json:"rewards"`
	BlockTime         uint64                             `json:"blockTime"`
	BlockHeight       uint64                             `json:"blockHeight"`
}
