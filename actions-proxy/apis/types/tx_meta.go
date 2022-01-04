package types

type TxMetaPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            TxMetaArgs             `json:"input"`
}

type TxMetaArgs struct {
	Address string `json:"address"`
	Limit   int    `json:"limit"`
	Before  string `json:"before"`
	Until   string `json:"until"`
}
