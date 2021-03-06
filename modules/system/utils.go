package system

import (
	"encoding/base64"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
)

// UpdateNonceAccount properly updates the statement of nonce inside the database
func UpdateNonceAccount(address string, currentSlot uint64, db db.SystemDb, client ClientProxy) error {
	if db.CheckNonceAccountLatest(address, currentSlot) {
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

	nonce, ok := parser.Parse(info.Value.Owner, bz).(parser.NonceAccount)
	if !ok {
		return db.DeleteNonceAccount(address)
	}

	return db.SaveNonceAccount(
		dbtypes.NewNonceAccountRow(
			address,
			info.Context.Slot,
			nonce.Authority.String(),
			nonce.BlockHash.String(),
			nonce.FeeCalculator.LamportsPerSignature,
		),
	)
}
