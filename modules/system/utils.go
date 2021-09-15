package system

import (
	"encoding/base64"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/solana/types"
)

// updateNonce properly stores the statement of nonce inside the database
func updateNonce(address string, db db.SystemDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveNonce(address, info.Context.Slot, "", "", types.FeeCalculator{}, "closed")
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	nonce, ok := accountParser.Parse(system.ProgramID, bz).(accountParser.NonceAccount)
	if !ok {
		return db.SaveNonce(address, info.Context.Slot, "", "", types.FeeCalculator{}, "closed")
	}

	return db.SaveNonce(address, info.Context.Slot, nonce.Authority.String(), nonce.BlockHash.String(), nonce.FeeCalculator, nonce.State.String())
}
