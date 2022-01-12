package types

import (
	"time"

	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/solana/parser/manager"
	"github.com/forbole/soljuno/solana/types"
)

// Validator contains the data of a single validator
type Validator struct {
	VotePubkey string
	NodePubKey string
}

// NewValidator allows to build a new Validator instance
func NewValidator(votePubkey string, nodePubKey string) Validator {
	return Validator{
		VotePubkey: votePubkey,
		NodePubKey: nodePubKey,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Block contains the data of a single chain block
type Block struct {
	Slot      uint64
	Height    uint64
	Hash      string
	Proposer  string
	Rewards   []clienttypes.Reward
	Timestamp time.Time
	Txs       []Tx
}

// NewBlock allows to build a new Block instance
func NewBlock(slot, height uint64, hash, proposer string, timestamp time.Time, txs []Tx) Block {
	return Block{
		Slot:      slot,
		Height:    height,
		Hash:      hash,
		Proposer:  proposer,
		Timestamp: timestamp,
		Txs:       txs,
	}
}

// NewBlockFromResult allows to build new Block instance from BlockResult
func NewBlockFromResult(parserManager manager.ParserManager, slot uint64, b clienttypes.BlockResult) Block {
	proposer := ""
	rewards := b.Rewards
	for _, reward := range rewards {
		if reward.RewardType == clienttypes.RewardFee {
			proposer = reward.Pubkey
			break
		}
	}

	var txs []Tx
	for _, txResult := range b.Transactions {
		txs = append(
			txs,
			NewTxFromTxResult(parserManager, slot, txResult),
		)
	}

	return Block{
		Slot:      slot,
		Height:    b.BlockHeight,
		Hash:      b.Blockhash,
		Rewards:   rewards,
		Proposer:  proposer,
		Timestamp: time.Unix(int64(b.BlockTime), 0),
		Txs:       txs,
	}
}

// Get account pubkeys from ids
func getAccounts(accountKeys []string, ids []uint8) []string {
	accounts := make([]string, len(ids))
	// Get account pubkey from id
	for i, id := range ids {
		accounts[i] = accountKeys[id]
	}
	return accounts
}

// -------------------------------------------------------------------------------------------------------------------

// Tx represents an already existing blockchain transaction
type Tx struct {
	Hash     string
	Slot     uint64
	Error    interface{}
	Fee      uint64
	Logs     []string
	Messages []Message

	Accounts          []string
	PostBalances      []uint64
	PostTokenBalances []clienttypes.TransactionTokenBalance
}

// NewTx allows to build a new Tx instance
func NewTx(
	hash string,
	slot uint64,
	err interface{},
	fee uint64,
	logs []string,
	msgs []Message,
	accounts []string,
	postBalances []uint64,
	postTokenBalances []clienttypes.TransactionTokenBalance,
) Tx {
	return Tx{
		Hash:     hash,
		Slot:     slot,
		Error:    err,
		Fee:      fee,
		Logs:     logs,
		Messages: msgs,

		Accounts:          accounts,
		PostBalances:      postBalances,
		PostTokenBalances: postTokenBalances,
	}
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return tx.Error == nil
}

func NewTxFromTxResult(parserManager manager.ParserManager, slot uint64, txResult clienttypes.EncodedTransactionWithStatusMeta) Tx {
	hash := txResult.Transaction.Signatures[0]

	var msgs []Message
	rawMsg := txResult.Transaction.Message
	accountKeys := rawMsg.AccountKeys

	// Put innerstructions to map in order to create msg after the main instruction
	var innerInstructionMap = make(map[uint8][]clienttypes.UiCompiledInstruction)
	for _, inner := range txResult.Meta.InnerInstructions {
		innerInstructionMap[inner.Index] = append(innerInstructionMap[inner.Index], inner.Instructions...)
	}

	for i, msg := range rawMsg.Instructions {
		var accounts []string
		innerIndex := 0
		accounts = getAccounts(accountKeys, msg.Accounts)
		programID := accountKeys[msg.ProgramIDIndex]
		parsed := parserManager.Parse(accounts, programID, msg.Data)
		msgs = append(msgs, NewMessage(hash, slot, i, innerIndex, accountKeys[msg.ProgramIDIndex], accounts, msg.Data, parsed))
		innerIndex++

		if inner, ok := innerInstructionMap[uint8(i)]; ok {
			for _, innerMsg := range inner {
				accounts = getAccounts(accountKeys, innerMsg.Accounts)
				programID := accountKeys[innerMsg.ProgramIDIndex]
				parsed := parserManager.Parse(accounts, programID, innerMsg.Data)
				msgs = append(msgs, NewMessage(hash, slot, i, innerIndex, accountKeys[innerMsg.ProgramIDIndex], accounts, innerMsg.Data, parsed))
				innerIndex++
			}
		}
	}
	return NewTx(
		hash,
		slot,
		txResult.Meta.Err,
		txResult.Meta.Fee,
		txResult.Meta.LogMessages,
		msgs,
		txResult.Transaction.Message.AccountKeys,
		txResult.Meta.PostBalances,
		txResult.Meta.PreTokenBalances,
	)
}

// -------------------------------------------------------------------------------------------------------------------

type Message struct {
	TxHash           string
	Slot             uint64
	Index            int
	InnerIndex       int
	Program          string
	InvolvedAccounts []string
	RawData          string
	Parsed           types.ParsedInstruction
}

func NewMessage(
	hash string,
	slot uint64,
	index int,
	innerIndex int,
	program string,
	involvedAccounts []string,
	rawData string,
	parsed types.ParsedInstruction,
) Message {
	return Message{
		TxHash:           hash,
		Slot:             slot,
		Index:            index,
		InnerIndex:       innerIndex,
		Program:          program,
		InvolvedAccounts: involvedAccounts,
		RawData:          rawData,
		Parsed:           parsed,
	}
}

type SanitizedMessage struct {
	Index            int                     `json:"index"`
	InnerIndex       int                     `json:"inner_index"`
	Program          string                  `json:"program"`
	InvolvedAccounts []string                `json:"involved_accounts"`
	RawData          string                  `json:"raw_data"`
	Parsed           types.ParsedInstruction `json:"parsed"`
}

func NewSanitizedMessages(msgs []Message) []SanitizedMessage {
	sanitizedMsgs := make([]SanitizedMessage, len(msgs))
	for i, msg := range msgs {
		sanitizedMsgs[i] = SanitizedMessage{
			Index:            msg.Index,
			InnerIndex:       msg.InnerIndex,
			Program:          msg.Program,
			InvolvedAccounts: msg.InvolvedAccounts,
			RawData:          msg.RawData,
			Parsed:           msg.Parsed,
		}
	}
	return sanitizedMsgs
}
