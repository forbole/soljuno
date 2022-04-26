package vote

import (
	"encoding/base64"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
)

// updateVoteAccount properly stores the statement of vote account inside the database
func updateVoteAccount(address string, currentSlot uint64, db db.VoteDb, client ClientProxy) error {
	if db.CheckValidatorLatest(address, currentSlot) {
		return nil
	}
	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}
	if info.Value == nil {
		return nil
	}
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	voteAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.VoteAccount)
	if !ok {
		return nil
	}
	return db.SaveValidator(
		dbtypes.NewVoteAccountRow(
			address,
			info.Context.Slot,
			voteAccount.Node.String(),
			voteAccount.Withdrawer.String(),
			voteAccount.Voters[0].Pubkey.String(),
			voteAccount.Commission,
		),
	)
}
