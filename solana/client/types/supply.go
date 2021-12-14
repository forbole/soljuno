package types

type SupplyWithContext struct {
	Context Context     `json:"context"`
	Value   SupplyValue `json:"value"`
}

type SupplyValue struct {
	Total                  uint64   `json:"total"`
	Circulating            uint64   `json:"circulating"`
	NonCirculating         uint64   `json:"nonCirculating"`
	NonCirculatingAccounts []string `json:"nonCirculatingAccounts"`
}

type SupplyConfig struct {
	ExcludeNonCirculatingAccountsList bool `json:"excludeNonCirculatingAccountsList,omitempty"`
}

func NewSupplyConfig(exclude bool) SupplyConfig {
	return SupplyConfig{
		ExcludeNonCirculatingAccountsList: exclude,
	}
}
