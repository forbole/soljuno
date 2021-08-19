package vote

import "github.com/forbole/soljuno/solana/types"

type InstructionID uint32

const (
	InitializeAccount InstructionID = iota
	Authorize
	Vote
	Withdraw
	UpdateValidatorIdentity
	UpdateCommission
	VoteSwitch
	AuthorizeChecked
)

type Instruction struct {
	ID   InstructionID
	Data interface{}
}

func (v *Instruction) Marshal([]byte) error {
	return nil
}

type VoteInstruction struct {
	Vote types.Vote
}
