package types

import "time"

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
	Hash      string
	Proposer  string
	Timestamp time.Time
}

// NewBlock allows to build a new Block instance
func NewBlock(slot uint64, hash, proposer string, timestamp time.Time) Block {
	return Block{
		Slot:      slot,
		Hash:      hash,
		Proposer:  proposer,
		Timestamp: timestamp,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Tx represents an already existing blockchain transaction
type Tx struct {
	Hash  string
	Slot  uint64
	Error bool
	Fee   int
	Logs  []string
}

// NewTx allows to build a new Tx instance
func NewTx() Tx {
	return Tx{}
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return !tx.Error
}

// -------------------------------------------------------------------------------------------------------------------

type Instruction struct {
	TxHash            string
	Index             int
	Program           string
	InnerInstructions []InnerInstruction
	Type              string
	Value             interface{}
}

type InnerInstruction struct {
	Program string
	Type    string
	Value   interface{}
}

func NewInstruction() Instruction {
	return Instruction{}
}

func NewInnerInstruction() InnerInstruction {
	return InnerInstruction{}
}
