package types

type EncodedConfirmedTransactionWithStatusMeta struct {
	Slot        uint64                  `json:"slot"`
	BlockTime   uint64                  `json:"blockTime"`
	Transaction UiTransaction           `json:"transaction"`
	Meta        UiTransactionStatusMeta `json:"meta"`
}

type GetSignaturesForAddressConfig struct {
	Limit  int    `json:"limit,omitempty"`
	Before string `json:"before,omitempty"`
	Until  string `json:"until,omitempty"`
}

type ConfirmedTransactionStatusWithSignature struct {
	Signature string      `json:"signature"`
	Slot      uint64      `json:"slot"`
	Err       interface{} `json:"err"`
	Memo      string      `json:"memo"`
	BlockTime uint64      `json:"blockTime"`
}

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
	LoadedAddresses   LoadedAddress             `json:"loadedAddresses"`
}

type LoadedAddress struct {
	Readonly []string `json:"readonly"`
	Writable []string `json:"writable"`
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
