package snapshot

import (
	"strconv"
	"strings"
)

type Account struct {
	Pubkey string
	Detail AccountDetail `yaml:"account"`
}

type AccountDetail struct {
	Balance string `yaml:"balance"`
	Owner   string `yaml:"owner"`
	Slot    uint64 `yaml:"slot"`
}

func (a Account) ToBalance() (string, uint64, error) {
	balanceStr := a.Detail.Balance
	balanceStr = strings.Replace(balanceStr, " SOL", "", 1)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	return a.Pubkey, ToLamports(balance), err
}

func ToLamports(balance float64) uint64 {
	return uint64(balance * 1_000_000_000)
}
