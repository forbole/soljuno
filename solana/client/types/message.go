package types

type UiRawMessage struct {
	Header          MessageHeader           `json:"header"`
	AccountKeys     []AccountKey            `json:"accountKeys"`
	RecentBlockHash string                  `json:"recentBlockHash"`
	Instructions    []UiCompiledInstruction `json:"instructions"`
}

type AccountKey struct {
	Pubkey   string `json:"pubkey"`
	Signer   bool   `json:"signer"`
	Source   string `json:"source"`
	Writable bool   `json:"writable"`
}

type MessageHeader struct {
	NumRequiredSignature        uint8 `json:"numRequiredSignature"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}
