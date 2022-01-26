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

func NewMsgRowFromMessage(
	msg types.Message,
) MsgRow {
	return MsgRow{
		TxHash:           msg.TxHash,
		Slot:             msg.Slot,
		Index:            msg.Index,
		InnerIndex:       msg.InnerIndex,
		Program:          msg.Program,
		InvolvedAccounts: msg.InvolvedAccounts,
		RawData:          msg.RawData,
		Type:             msg.Parsed.Type,
		Value:            msg.Parsed.GetValueJSON(),
		PartitionId:      int(msg.Slot / 1000),
	}
}
