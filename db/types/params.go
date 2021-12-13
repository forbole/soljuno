package types

type EpochScheduleParamRow struct {
	Epoch            uint64 `db:"epoch"`
	SlotsPerEpoch    uint64 `db:"slotsPerEpoch"`
	FirstNormalEpoch uint64 `db:"firstNormalEpoch"`
	FirstNormalSlot  uint64 `db:"firstNormalSlot"`
}

func NewEpochScheduleParamRow(
	epoch uint64,
	slotsPerEpoch uint64,
	firstNormalEpoch uint64,
	firstNormalSlot uint64,
) EpochScheduleParamRow {
	return EpochScheduleParamRow{
		epoch,
		slotsPerEpoch,
		firstNormalEpoch,
		firstNormalSlot,
	}
}

type InflationGovernorParamRow struct {
	Epoch      uint64  `db:"epoch"`
	Total      float64 `db:"total"`
	Foundation float64 `db:"foundation"`
	Validator  float64 `db:"validator"`
}

func NewInflationGovernParamRow(
	epoch uint64,
	total float64,
	foundation float64,
	validator float64,
) InflationGovernorParamRow {
	return InflationGovernorParamRow{
		epoch,
		total,
		foundation,
		validator,
	}
}
