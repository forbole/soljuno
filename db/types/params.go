package types

type EpochScheduleParamRow struct {
	Epoch            uint64 `db:"epoch"`
	SlotsPerEpoch    uint64 `db:"slotsPerEpoch"`
	FirstNormalEpoch uint64 `db:"firstNormalEpoch"`
	FirstNormalSlot  uint64 `db:"firstNormalSlot"`
	Warmup           bool   `db:"warmup"`
}

func NewEpochScheduleParamRow(
	epoch uint64,
	slotsPerEpoch uint64,
	firstNormalEpoch uint64,
	firstNormalSlot uint64,
	warmup bool,
) EpochScheduleParamRow {
	return EpochScheduleParamRow{
		epoch,
		slotsPerEpoch,
		firstNormalEpoch,
		firstNormalSlot,
		warmup,
	}
}

type InflationGovernorParamRow struct {
	Epoch              uint64  `db:"epoch"`
	Initial            float64 `db:"initial"`
	Terminal           float64 `db:"terminal"`
	Taper              float64 `db:"taper"`
	Foundation         float64 `db:"foundation"`
	FoundationTerminal float64 `db:"foundation_terminal"`
}

func NewInflationGovernParamRow(
	epoch uint64,
	initial float64,
	terminal float64,
	taper float64,
	foundation float64,
	foundationTerminal float64,
) InflationGovernorParamRow {
	return InflationGovernorParamRow{
		epoch,
		initial,
		terminal,
		taper,
		foundation,
		foundationTerminal,
	}
}
