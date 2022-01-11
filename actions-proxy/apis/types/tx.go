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
	var res TxResponse
	res.Hash = tx.Hash
	res.Slot = tx.Slot
	res.Error = !tx.Successful()
	res.Fee = tx.Fee
	res.Logs = tx.Logs
	res.Accounts = tx.Accounts
	for _, msg := range tx.Messages {
		res.Messages = append(res.Messages, MsgResponse{
			msg.Index,
			msg.InnerIndex,
			msg.Program,
			msg.InvolvedAccounts,
			msg.RawData,
			ParsedData{
				msg.Parsed.Type,
				msg.Parsed.Data,
			},
		})
	}
	return res
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
