package bank

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

func HandleBlock(block types.Block, db db.BankDb) error {
	var accounts []string
	var balances []uint64
	for _, reward := range block.Rewards {
		accounts = append(accounts, reward.Pubkey)
		balances = append(balances, reward.PostBalance)
	}
	return db.SaveAccountBalances(block.Slot, accounts, balances)
}
