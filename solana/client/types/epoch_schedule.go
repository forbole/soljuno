package types

type EpochSchedule struct {
	SlotsPerEpoch            uint64 `json:"slotsPerEpoch"`
	LeaderScheduleSlotOffset uint64 `json:"leaderScheduleSlotOffset"`
	Warmup                   bool   `json:"warmup"`
	FirstNormalEpoch         uint64 `json:"firstNormalEpoch"`
	FirstNormalSlot          uint64 `json:"firstNormalSlot"`
}
