package types

import (
	"encoding/base64"

	accountParser "github.com/forbole/soljuno/solana/account"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type AccountInfoPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AccountInfoArgs        `json:"input"`
}

type AccountInfoArgs struct {
	Address string `json:"address"`
}

type AccountInfoResponse struct {
	Data       [2]string   `json:"data"`
	Executable bool        `json:"executable"`
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"program_owner"`
	RentEpoch  uint64      `json:"rentepoch"`
	Parsed     interface{} `json:"parsed"`
}

func NewAccountInfoResponse(info clienttypes.AccountInfo) (AccountInfoResponse, error) {
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return AccountInfoResponse{}, err
	}
	parsed := accountParser.Parse(info.Value.Owner, bz)
	return AccountInfoResponse{
		Data:       info.Value.Data,
		Executable: info.Value.Executable,
		Lamports:   info.Value.Lamports,
		Owner:      info.Value.Owner,
		RentEpoch:  info.Value.RentEpoch,
		Parsed:     parsed,
	}, nil
}
