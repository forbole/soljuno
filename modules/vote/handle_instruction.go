package vote

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/vote"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the vote module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initialize":
		return handleInitialize(instruction, tx, db)
	case "authorize":
		return handleAuthorize(instruction, tx, db, client)
	case "withdraw":
		return handleWithdraw(instruction, tx, db, client)
	case "updateValidatorIdentity":
		return handleUpdateValidatorIdentity(instruction, tx, db, client)
	case "updateCommission":
		return handleUpdateCommission(instruction, tx, db, client)
	case "authorizeChecked":
		return handleAuthorizeChecked(instruction, tx, db, client)
	}
	return nil
}

// handleInitialize handles a instruction of Initialize
func handleInitialize(instruction types.Instruction, tx types.Tx, db db.VoteDb) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", instruction.Parsed.Type)

	}
	return db.SaveValidator(parsed.VoteAccount, tx.Slot, parsed.Node, parsed.AuthorizedVoter, parsed.AuthorizedWithdrawer, parsed.Commission)
}

// handleAuthorize handles a instruction of Authorize
func handleAuthorize(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", instruction.Parsed.Type)

	}
	return updateVoteAccount(parsed.VoteAccount, tx.Slot, db, client)
}

// handleWithdraw handles a instruction of Withdraw
func handleWithdraw(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", instruction.Parsed.Type)

	}
	return updateVoteAccount(parsed.VoteAccount, tx.Slot, db, client)
}

// handleUpdateValidatorIdentity handles a instruction of UpdateValidatorIdentity
func handleUpdateValidatorIdentity(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedUpdateValidatorIdentity)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateValidatorIdentity", instruction.Parsed.Type)

	}
	return updateVoteAccount(parsed.VoteAccount, tx.Slot, db, client)
}

// handleUpdateCommission handles a instruction of UpdateCommission
func handleUpdateCommission(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedUpdateCommission)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "updateCommission", instruction.Parsed.Type)

	}
	return updateVoteAccount(parsed.VoteAccount, tx.Slot, db, client)
}

// handleAuthorizeChecked handles a instruction of AuthorizeChecked
func handleAuthorizeChecked(instruction types.Instruction, tx types.Tx, db db.VoteDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(vote.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", instruction.Parsed.Type)

	}
	return updateVoteAccount(parsed.VoteAccount, tx.Slot, db, client)
}
