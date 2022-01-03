package types

type InflationGovernor struct {
	Initial            float64 `json:"initial"`
	Terminal           float64 `json:"terminal"`
	Taper              float64 `json:"taper"`
	Foundation         float64 `json:"foundation"`
	FoundationTerminal float64 `json:"foundationTerm"`
}
