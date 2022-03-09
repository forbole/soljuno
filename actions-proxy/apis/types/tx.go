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

type TxResponse struct {
	Hash         string                `json:"hash"`
	Slot         uint64                `json:"slot"`
	Error        bool                  `json:"error"`
	Fee          uint64                `json:"fee"`
	Logs         []string              `json:"logs"`
	Instructions []InstructionResponse `json:"instructions"`

	Accounts []string `json:"accounts"`
}

type InstructionResponse struct {
	Index            int        `json:"index"`
	InnerIndex       int        `json:"inner_index"`
	Program          string     `json:"program"`
	InvolvedAccounts []string   `json:"involved_accounts"`
	RawData          string     `json:"raw_data"`
	Parsed           ParsedData `json:"parsed"`
}

type ParsedData struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func NewTxResponse(tx types.Tx) TxResponse {
	var res TxResponse
	res.Hash = tx.Hash
	res.Slot = tx.Slot
	res.Error = !tx.Successful()
	res.Fee = tx.Fee
	res.Logs = tx.Logs
	res.Accounts = tx.Accounts
	res.Instructions = make([]InstructionResponse, len(tx.Instructions))
	for i, msg := range tx.Instructions {
		res.Instructions[i] = InstructionResponse{
			msg.Index,
			msg.InnerIndex,
			msg.Program,
			msg.InvolvedAccounts,
			msg.RawData,
			ParsedData{
				msg.Parsed.Type,
				msg.Parsed.Value,
			},
		}
	}
	return res
}
