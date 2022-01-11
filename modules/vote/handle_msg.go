package vote

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/vote"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the vote module
func HandleMsg(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	switch msg.Parsed.Type {
	case "initialize":
		return handleMsgInitialize(msg, tx, db)
	case "authorize":
		return handleMsgAuthorize(msg, tx, db, client)
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
	instruction, ok := msg.Parsed.Value.(vote.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", msg.Parsed.Type)

	}
	return db.SaveValidator(instruction.VoteAccount, tx.Slot, instruction.Node, instruction.AuthorizedVoter, instruction.AuthorizedWithdrawer, instruction.Commission)
}

// handleMsgAuthorize handles a MsgAuthorize
func handleMsgAuthorize(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(vote.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", msg.Parsed.Type)

	}
	return updateVoteAccount(instruction.VoteAccount, tx.Slot, db, client)
}

// handleMsgWithdraw handles a MsgWithdraw
func handleMsgWithdraw(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(vote.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", msg.Parsed.Type)

	}
	return updateVoteAccount(instruction.VoteAccount, tx.Slot, db, client)
}

// handleMsgUpdateValidatorIdentity handles a MsgUpdateValidatorIdentity
func handleMsgUpdateValidatorIdentity(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(vote.ParsedUpdateValidatorIdentity)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateValidatorIdentity", msg.Parsed.Type)

	}
	return updateVoteAccount(instruction.VoteAccount, tx.Slot, db, client)
}

// handleMsgUpdateCommission handles a MsgUpdateCommission
func handleMsgUpdateCommission(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(vote.ParsedUpdateCommission)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateCommission", msg.Parsed.Type)

	}
	return updateVoteAccount(instruction.VoteAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeChecked handles a MsgAuthorizeChecked
func handleMsgAuthorizeChecked(msg types.Message, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(vote.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", msg.Parsed.Type)

	}
	return updateVoteAccount(instruction.VoteAccount, tx.Slot, db, client)
}
