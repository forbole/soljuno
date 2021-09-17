package vote

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/vote"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the vote module
func HandleMsg(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "initialize":
		return handleMsgInitialize(msg, tx, db)
	case "authorize":
		return handleMsgAuthorize(msg, tx, db, client)
	case "withdraw":
		return handleMsgWithdraw(msg, tx, db, client)
	case "updateValidatorIdentity":
		return handleMsgUpdateValidatorIdentity(msg, tx, db, client)
	case "updateCommission":
		return handleMsgUpdateCommission(msg, tx, db, client)
	case "authorizeChecked":
		return handleMsgAuthorizeChecked(msg, tx, db, client)
	}
	return nil
}

// handleMsgInitialize handles a MsgInitialize
func handleMsgInitialize(msg types.Message, tx types.Tx, db db.VoteDb) error {
	instruction, ok := msg.Value.Data().(vote.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", msg.Value.Type())

	}
	return db.SaveVoteAccount(instruction.VoteAccount, tx.Slot, instruction.Node, instruction.AuthorizedVoter, instruction.AuthorizedWithdrawer, instruction.Commission)
}

// handleMsgAuthorize handles a MsgAuthorize
func handleMsgAuthorize(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(vote.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", msg.Value.Type())

	}
	return updateVoteAccount(instruction.VoteAccount, db, client)
}

// handleMsgWithdraw handles a MsgWithdraw
func handleMsgWithdraw(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(vote.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", msg.Value.Type())

	}
	return updateVoteAccount(instruction.VoteAccount, db, client)
}

// handleMsgUpdateValidatorIdentity handles a MsgUpdateValidatorIdentity
func handleMsgUpdateValidatorIdentity(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(vote.ParsedUpdateValidatorIdentity)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateValidatorIdentity", msg.Value.Type())

	}
	return updateVoteAccount(instruction.VoteAccount, db, client)
}

// handleMsgUpdateCommission handles a MsgUpdateCommission
func handleMsgUpdateCommission(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(vote.ParsedUpdateCommission)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateCommission", msg.Value.Type())

	}
	return updateVoteAccount(instruction.VoteAccount, db, client)
}

// handleMsgAuthorizeChecked handles a MsgAuthorizeChecked
func handleMsgAuthorizeChecked(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(vote.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", msg.Value.Type())

	}
	return updateVoteAccount(instruction.VoteAccount, db, client)
}
