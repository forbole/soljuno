package associated_token_account

type ParsedCreate struct {
	System  string `json:"system"`
	Account string `json:"account"`
	Owner   string `json:"owner"`
	Mint    string `json:"mint"`
}

func NewParsedCreate(
	accounts []string,
) ParsedCreate {
	return ParsedCreate{
		System:  accounts[0],
		Account: accounts[1],
		Owner:   accounts[2],
		Mint:    accounts[3],
	}
}
