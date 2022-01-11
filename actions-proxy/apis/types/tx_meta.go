package types

import clienttypes "github.com/forbole/soljuno/solana/client/types"

type TxByAddressPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            TxByAddressArgs        `json:"input"`
}

type TxByAddressArgs struct {
	Address string            `json:"address"`
	Config  TxByAddressConfig `json:"config"`
}

type TxByAddressConfig struct {
	Limit  int    `json:"limit"`
	Before string `json:"before"`
	Until  string `json:"until"`
}

type TxMetaRespoonse struct {
	Hash      string `json:"hash"`
	Slot      uint64 `json:"slot"`
	Error     bool   `json:"error"`
	Memo      string `json:"memo"`
	BlockTime uint64 `json:"block_time"`
}

func NewTxMetasResponse(metas []clienttypes.ConfirmedTransactionStatusWithSignature) []TxMetaRespoonse {
	res := make([]TxMetaRespoonse, len(metas))
	for i, meta := range metas {
		res[i] = TxMetaRespoonse{
			Hash:      meta.Signature,
			Slot:      meta.Slot,
			Error:     meta.Err != nil,
			Memo:      meta.Memo,
			BlockTime: meta.BlockTime,
		}
	}
	return res
}
