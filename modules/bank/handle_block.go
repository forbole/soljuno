package bank

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

func HandleBlock(block types.Block, db db.BankDb) error {
	var accounts []string
	var balances []uint64
	m := make(map[string]int)
	for i, reward := range block.Rewards {
		if preId, ok := m[reward.Pubkey]; ok {
			balances[preId] = reward.PostBalance
			m[reward.Pubkey] = i
			continue
		}
		accounts = append(accounts, reward.Pubkey)
		balances = append(balances, reward.PostBalance)
		m[reward.Pubkey] = i
	}
	return db.SaveAccountBalances(block.Slot, accounts, balances)
}
