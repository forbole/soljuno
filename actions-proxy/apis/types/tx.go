package types

import (
	"github.com/forbole/soljuno/types"
)

type TxPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            TxArgs                 `json:"input"`
}

type TxArgs struct {
	Signature string `json:"signature"`
}

type TxResponse struct {
	Signature    string                `json:"signature"`
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
	res.Signature = tx.Signature
	res.Slot = tx.Slot
	res.Error = !tx.Successful()
	res.Fee = tx.Fee
	res.Logs = tx.Logs
	res.Accounts = tx.Accounts
	res.Instructions = make([]InstructionResponse, len(tx.Instructions))
	for i, instruction := range tx.Instructions {
		res.Instructions[i] = InstructionResponse{
			instruction.Index,
			instruction.InnerIndex,
			instruction.Program,
			instruction.InvolvedAccounts,
			instruction.RawData,
			ParsedData{
				instruction.Parsed.Type,
				instruction.Parsed.Value,
			},
		}
	}
	return res
}
