package types

import "github.com/forbole/soljuno/types"

type InstructionRow struct {
	TxSignature      string      `db:"tx_signature"`
	Slot             uint64      `db:"slot"`
	Index            int         `db:"index"`
	InnerIndex       int         `db:"inner_index"`
	Program          string      `db:"program"`
	InvolvedAccounts interface{} `db:"involved_accounts"`
	RawData          string      `db:"raw_data"`
	Type             string      `db:"type"`
	Value            interface{} `db:"value"`
	PartitionId      int         `db:"partition_id"`
}

func NewInstructionRow(
	txSignature string, slot uint64, index int, innerIndex int, program string, involvedAccounts []string, rawData string, typ string, value interface{},
) InstructionRow {
	return InstructionRow{
		TxSignature:      txSignature,
		Slot:             slot,
		Index:            index,
		InnerIndex:       innerIndex,
		Program:          program,
		InvolvedAccounts: involvedAccounts,
		RawData:          rawData,
		Type:             typ,
		Value:            value,
		PartitionId:      int(slot / 1000),
	}
}

func NewInstructionRowFromInstruction(
	instruction types.Instruction,
) InstructionRow {
	return NewInstructionRow(
		instruction.TxSignature,
		instruction.Slot,
		instruction.Index,
		instruction.InnerIndex,
		instruction.Program,
		instruction.InvolvedAccounts,
		instruction.RawData,
		instruction.Parsed.Type,
		instruction.Parsed.GetValueJSON(),
	)
}
