package types

type EncodedTransactionWithStatusMeta struct {
	Transaction UiTransaction           `json:"transaction"`
	Meta        UiTransactionStatusMeta `json:"meta"`
}

type UiTransactionStatusMeta struct {
	Err               interface{}               `json:"err"`
	Fee               uint64                    `json:"fee"`
	PreBalances       []uint64                  `json:"preBalances"`
	PostBalances      []uint64                  `json:"postBalances"`
	InnerInstructions []UiInnerInstruction      `json:"innerInstructions"`
	LogMessages       []string                  `json:"logMessages"`
	PreTokenBalances  []TransactionTokenBalance `json:"preTokenBalances"`
	PostTokenBalances []TransactionTokenBalance `json:"postTokenBalances"`
	Rewards           []Reward                  `json:"rewards"`
}

type UiTransaction struct {
	Signatures []string     `json:"signatures"`
	Message    UiRawMessage `json:"message"`
}

type TransactionTokenBalance struct {
	AccountIndex  uint          `json:"accountIndex"`
	Mint          string        `json:"mint"`
	UiTokenAmount UiTokenAmount `json:"uiTokenAmount"`
}

type UiTokenAmount struct {
	UiAmount       float64 `json:"uiAmount"`
	Decimals       uint8   `json:"decimals"`
	Amount         string  `json:"amount"`
	UiAmountString string  `json:"uiAmountString,omitempty"`
}
