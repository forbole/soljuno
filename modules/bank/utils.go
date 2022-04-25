package bank

import (
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
)

func getAccountBalaceEntries(block types.Block) []AccountBalanceEntry {
	var accounts []string
	var balances []uint64
	accountMap := make(map[string]int)
	accountIndex := 0
	// Gather balances from rewards
	for _, reward := range block.Rewards {
		if index, ok := accountMap[reward.Pubkey]; ok {
			balances[index] = reward.PostBalance
			continue
		}
		accounts = append(accounts, reward.Pubkey)
		balances = append(balances, reward.PostBalance)
		accountMap[reward.Pubkey] = accountIndex
		accountIndex++
	}

	for _, tx := range block.Txs {
		// Gather balances from tx
		for i, pubkey := range tx.Accounts {
			if index, ok := accountMap[pubkey]; ok {
				balances[index] = tx.PostBalances[i]
				continue
			}
			accounts = append(accounts, pubkey)
			balances = append(balances, tx.PostBalances[i])
			accountMap[pubkey] = accountIndex
			accountIndex++
		}
	}
	return NewAccountBalanceEntries(block.Slot, accounts, balances)
}

func getTokenAccountBalaceEntries(block types.Block) []TokenAccountBalanceEntry {
	var tokenAccounts []string
	var tokenBalances []clienttypes.TransactionTokenBalance
	tokenAccountMap := make(map[string]int)
	tokenAccountIndex := 0
	for _, tx := range block.Txs {
		// Gather token balances from tx
		for _, balance := range tx.PostTokenBalances {
			pubkey := tx.Accounts[balance.AccountIndex]
			if index, ok := tokenAccountMap[pubkey]; ok {
				balance.AccountIndex = uint(index)
				tokenBalances[index] = balance
				continue
			}
			balance.AccountIndex = uint(tokenAccountIndex)
			tokenBalances = append(tokenBalances, balance)
			tokenAccounts = append(tokenAccounts, pubkey)
			tokenAccountMap[pubkey] = tokenAccountIndex
			tokenAccountIndex++
		}
	}
	return NewTokenAccountBalanceEntries(block.Slot, tokenAccounts, tokenBalances)
}
