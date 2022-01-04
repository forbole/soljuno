package types

import (
	"github.com/forbole/soljuno/types"
)

type TxPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            TxArgs                 `json:"input"`
}

type TxArgs struct {
	Hash string `json:"hash"`
}

type ParsedData struct {
	ParsedType string      `json:"type"`
	Value      interface{} `json:"value"`
}

func NewTxResponse(tx types.Tx) TxResponse {
	var txResponse TxResponse
	txResponse.Hash = tx.Hash
	txResponse.Slot = tx.Slot
	txResponse.Error = !tx.Successful()
	txResponse.Fee = tx.Fee
	txResponse.Accounts = tx.Accounts
	for _, msg := range tx.Messages {
		txResponse.Messages = append(txResponse.Messages, MsgResponse{
			msg.Index,
			msg.InnerIndex,
			msg.Program,
			msg.InvolvedAccounts,
			msg.RawData,
			ParsedData{
				msg.Parsed.Type(),
				msg.Parsed.Data(),
			},
		})
	}
	return txResponse
}

type TxResponse struct {
	Hash     string        `json:"hash"`
	Slot     uint64        `json:"slot"`
	Error    bool          `json:"error"`
	Fee      uint64        `json:"fee"`
	Logs     []string      `json:"logs"`
	Messages []MsgResponse `json:"messages"`

	Accounts []string `json:"accounts"`
}

type MsgResponse struct {
	Index            int        `json:"index"`
	InnerIndex       int        `json:"innerIndex"`
	Program          string     `json:"program"`
	InvolvedAccounts []string   `json:"involvedAccounts"`
	RawData          string     `json:"rawData"`
	Parsed           ParsedData `json:"parsed"`
}
