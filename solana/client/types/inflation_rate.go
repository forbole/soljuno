package types

type InflationRate struct {
	Total      float64 `json:"total"`
	Validator  float64 `json:"validator"`
	Foundation float64 `json:"foundation"`
	Epoch      uint64  `json:"epoch"`
}
