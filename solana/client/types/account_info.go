package types

type AccountInfoOption struct {
	Encoding string `json:"encoding,omitempty"`
}

func NewAccountInfoOption(encoding string) AccountInfoOption {
	return AccountInfoOption{
		Encoding: encoding,
	}
}

type AccountInfo struct {
	Context Context       `json:"context"`
	Value   *AccountValue `json:"value"`
}

type AccountValue struct {
	Data       [2]string `json:"data"`
	Executable bool      `json:"executable"`
	Lamports   uint64    `json:"lamports"`
	Owner      string    `json:"owner"`
	RentEpoch  uint64    `json:"rentepoch"`
}
