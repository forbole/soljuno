package vote

import (
	"encoding/base64"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/program/vote"
)

func updateVoteAccount(address string, db db.VoteDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
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

	_, ok := accountParser.Parse(vote.ProgramID, bz).(accountParser.VoteAccount)
	if !ok {
		return nil
	}
	return nil
}
