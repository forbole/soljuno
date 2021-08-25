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
	Error interface{}
	Fee   int
	Logs  []string
}

// NewTx allows to build a new Tx instance
func NewTx(hash string, slot uint64, err interface{}, fee int, logs []string) Tx {
	return Tx{
		Hash:  hash,
		Slot:  slot,
		Error: err,
		Fee:   fee,
		Logs:  logs,
	}
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return tx.Error == nil
}

// -------------------------------------------------------------------------------------------------------------------

type Instruction struct {
	TxHash            string
	Index             int
	Program           string
	InvolvedAccounts  []string
	InnerInstructions []InnerInstruction
	Type              string
	Value             interface{}
}

type InnerInstruction struct {
	Program string      `json:"program"`
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
}

func NewInstruction(hash string, index int, program string, involvedAccounts []string, innerInstructions []InnerInstruction, typ string, value interface{}) Instruction {
	return Instruction{
		TxHash:            hash,
		Index:             index,
		Program:           program,
		InvolvedAccounts:  involvedAccounts,
		InnerInstructions: innerInstructions,
		Type:              typ,
		Value:             value,
	}
}

func NewInnerInstruction(program, typ string, value interface{}) InnerInstruction {
	return InnerInstruction{
		Program: program,
		Type:    typ,
		Value:   value,
	}
}
