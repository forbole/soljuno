package types

import "github.com/forbole/soljuno/types"

type MsgRow struct {
	TxHash           string      `db:"transaction_hash"`
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

func NewMsgRow(
	txHash string, slot uint64, index int, innerIndex int, program string, involvedAccounts []string, rawData string, typ string, value interface{},
) MsgRow {
	return MsgRow{
		TxHash:           txHash,
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

func NewMsgRowFromMessage(
	msg types.Message,
) MsgRow {
	return NewMsgRow(
		msg.TxHash,
		msg.Slot,
		msg.Index,
		msg.InnerIndex,
		msg.Program,
		msg.InvolvedAccounts,
		msg.RawData,
		msg.Parsed.Type,
		msg.Parsed.GetValueJSON(),
	)
}
