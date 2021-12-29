package system

import (
	"encoding/base64"

	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/client"
)

// updateNonce properly stores the statement of nonce inside the database
func updateNonce(address string, currentSlot uint64, db db.SystemDb, client client.ClientProxy) error {
	if !db.CheckNonceAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteNonceAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	nonce, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.NonceAccount)
	if !ok {
		return db.DeleteNonceAccount(address)
	}

	return db.SaveNonceAccount(
		address,
		info.Context.Slot,
		nonce.Authority.String(),
		nonce.BlockHash.String(),
		nonce.FeeCalculator.LamportsPerSignature,
	)
}
