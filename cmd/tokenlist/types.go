package tokenlist

import "time"

type TokenList struct {
	Name      string    `json:"name"`
	LogoURI   string    `json:"logoURI"`
	Keywords  []string  `json:"keywords"`
	Timestamp time.Time `json:"timestamp"`
	Tokens    []Token   `json:"tokens"`
}

type Token struct {
	ChainID    int        `json:"chainId"`
	Address    string     `json:"address"`
	Symbol     string     `json:"symbol"`
	Name       string     `json:"name"`
	Decimals   int        `json:"decimals"`
	LogoURI    string     `json:"logoURI"`
	Tags       []string   `json:"tags"`
	Extensions Extensions `json:"extensions"`
}

type Extensions struct {
	CoingeckoID string `json:"coingeckoId"`
	Description string `json:"description"`
	Website     string `json:"website"`
}
