package types

import clienttypes "github.com/forbole/soljuno/solana/client/types"

type InflationRateResponse struct {
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
	Foundation float64 `json:"foundation"`
	Epoch      uint64  `json:"epoch"`
}

func NewInflationRateResponse(inflation clienttypes.InflationRate) InflationRateResponse {
	return InflationRateResponse{
		Total:      inflation.Total,
		Validator:  inflation.Validator,
		Foundation: inflation.Foundation,
		Epoch:      inflation.Epoch,
	}
}

// --------------------------------------------------------------

type InflationGovernorResponse struct {
	Initial            float64 `json:"initial"`
	Terminal           float64 `json:"terminal"`
	Taper              float64 `json:"taper"`
	Foundation         float64 `json:"foundation"`
	FoundationTerminal float64 `json:"foundation_terminal"`
}

func NewInflationGovernorResponse(governor clienttypes.InflationGovernor) InflationGovernorResponse {
	return InflationGovernorResponse{
		Initial:            governor.Initial,
		Terminal:           governor.Terminal,
		Taper:              governor.Taper,
		Foundation:         governor.Foundation,
		FoundationTerminal: governor.FoundationTerminal,
	}
}
