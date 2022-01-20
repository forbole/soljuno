package bank

import (
	"strconv"

	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type AccountBalanceEntry struct {
	Slot    uint64
	Address string
	Balance uint64
}

func NewAccountBalanceEntry(slot uint64, address string, balance uint64) AccountBalanceEntry {
	return AccountBalanceEntry{
		slot,
		address,
		balance,
	}
}

func NewAccountBalanceEntries(slot uint64, addresses []string, balances []uint64) []AccountBalanceEntry {
	var entries []AccountBalanceEntry
	for i, address := range addresses {
		entries = append(entries, NewAccountBalanceEntry(slot, address, balances[i]))
	}
	return entries
}

func MergeAccountBalanceEntries(oldEntries, newEntries []AccountBalanceEntry) []AccountBalanceEntry {
	accountMap := make(map[string]int)
	for i, entry := range oldEntries {
		accountMap[entry.Address] = i
	}

	for _, entry := range newEntries {
		if i, exist := accountMap[entry.Address]; exist {
			if entry.Slot >= oldEntries[i].Slot {
				oldEntries[i].Slot = entry.Slot
				oldEntries[i].Balance = entry.Balance
			}
			continue
		}
		oldEntries = append(oldEntries, entry)
		accountMap[entry.Address] = len(oldEntries) - 1
	}
	return oldEntries
}

func EntriesToBalances(entries []AccountBalanceEntry) (uint64, []string, []uint64) {
	var slot uint64
	var accounts []string
	var balances []uint64
	for _, entry := range entries {
		accounts = append(accounts, entry.Address)
		balances = append(balances, entry.Balance)
		if entry.Slot > slot {
			slot = entry.Slot
		}
	}
	return slot, accounts, balances
}

// ----------------------------------------------------------------

type TokenAccountBalanceEntry struct {
	Slot    uint64
	Address string
	Balance uint64
}

func NewTokenAccountBalanceEntry(slot uint64, address string, balance uint64) TokenAccountBalanceEntry {
	return TokenAccountBalanceEntry{
		slot,
		address,
		balance,
	}
}

func NewTokenAccountBalanceEntries(slot uint64, addresses []string, balances []clienttypes.TransactionTokenBalance) []TokenAccountBalanceEntry {
	var entries []TokenAccountBalanceEntry
	for i, address := range addresses {
		bal, _ := strconv.ParseUint(balances[i].UiTokenAmount.Amount, 10, 64)
		entries = append(entries, NewTokenAccountBalanceEntry(slot, address, bal))
	}
	return entries
}

func MergeTokenAccountBalanceEntries(oldEntries, newEntries []TokenAccountBalanceEntry) []TokenAccountBalanceEntry {
	accountMap := make(map[string]int)
	for i, entry := range oldEntries {
		accountMap[entry.Address] = i
	}

	for _, entry := range newEntries {
		if i, exist := accountMap[entry.Address]; exist {
			if entry.Slot >= oldEntries[i].Slot {
				oldEntries[i].Balance = entry.Balance
			}
			continue
		}
		accountMap[entry.Address] = len(oldEntries)
		oldEntries = append(oldEntries, entry)
	}
	return oldEntries
}

func EntriesToTokenBalances(entries []TokenAccountBalanceEntry) (uint64, []string, []uint64) {
	var slot uint64
	var accounts []string
	var balances []uint64
	for _, entry := range entries {
		accounts = append(accounts, entry.Address)
		balances = append(balances, entry.Balance)
		if entry.Slot > slot {
			slot = entry.Slot
		}
	}
	return slot, accounts, balances
}
