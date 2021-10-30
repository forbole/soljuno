package vote

import (
	"encoding/base64"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
)

// updateVoteAccount properly stores the statement of vote account inside the database
func updateVoteAccount(address string, currentSlot uint64, db db.VoteDb, client client.Proxy) error {
	if !db.CheckVoteAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}
	if info.Value == nil {
		return db.SaveVoteAccount(address, info.Context.Slot, "", "", "", 0)
	}
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	voteAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.VoteAccount)
	if !ok {
		return db.SaveVoteAccount(address, info.Context.Slot, "", "", "", 0)
	}
	return db.SaveVoteAccount(
		address,
		info.Context.Slot,
		voteAccount.Node.String(),
		voteAccount.Withdrawer.String(),
		voteAccount.Voter[0].Pubkey.String(),
		voteAccount.Commission,
	)
}
