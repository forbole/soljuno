package types

import clienttypes "github.com/forbole/soljuno/solana/client/types"

type EpochInfoResponse struct {
	AbsoluteSlot     uint64 `json:"absolute_slot"`
	BlockHeight      uint64 `json:"block_height"`
	Epoch            uint64 `json:"epoch"`
	SlotIndex        uint64 `json:"slot_index"`
	SlotsInEpoch     uint64 `json:"slots_in_epoch"`
	TransactionCount uint64 `json:"transaction_count"`
}

func NewEpochInfoResponse(info clienttypes.EpochInfo) EpochInfoResponse {
	return EpochInfoResponse{
		AbsoluteSlot:     info.AbsoluteSlot,
		BlockHeight:      info.BlockHeight,
		Epoch:            info.Epoch,
		SlotIndex:        info.SlotIndex,
		SlotsInEpoch:     info.SlotsInEpoch,
		TransactionCount: info.TransactionCount,
	}
}

// ----------------------------------------------------------------------------

type EpochScheduleResponse struct {
	SlotsPerEpoch            uint64 `json:"slots_per_epoch"`
	LeaderScheduleSlotOffset uint64 `json:"leader_schedule_slot_offset"`
	Warmup                   bool   `json:"warmup"`
	FirstNormalEpoch         uint64 `json:"first_normal_epoch"`
	FirstNormalSlot          uint64 `json:"first_normal_slot"`
}

func NewEpochScheduleResponse(schedule clienttypes.EpochSchedule) EpochScheduleResponse {
	return EpochScheduleResponse{
		SlotsPerEpoch:            schedule.SlotsPerEpoch,
		LeaderScheduleSlotOffset: schedule.LeaderScheduleSlotOffset,
		Warmup:                   schedule.Warmup,
		FirstNormalEpoch:         schedule.FirstNormalEpoch,
		FirstNormalSlot:          schedule.FirstNormalSlot,
	}
}
