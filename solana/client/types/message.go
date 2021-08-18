package types

type UiRawMessage struct {
	Header          MessageHeader           `json:"header"`
	AccountKeys     []string                `json:"accountKeys"`
	RecentBlockHash string                  `json:"recentBlockHash"`
	Instructions    []UiCompiledInstruction `json:"instructions"`
}

type MessageHeader struct {
	NumRequiredSignature        uint8 `json:"numRequiredSignature"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}
