package bank_test

import (
	"testing"

	"github.com/forbole/soljuno/modules/bank"
	"github.com/stretchr/testify/require"
)

func TestMergeAccountBalanceEntries(t *testing.T) {
	oldEntries := []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(0, "address1", 1),
		bank.NewAccountBalanceEntry(0, "address2", 1),
	}
	newEntries := []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(1, "address1", 2),
		bank.NewAccountBalanceEntry(1, "address3", 1),
	}
	expected := []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(1, "address1", 2),
		bank.NewAccountBalanceEntry(0, "address2", 1),
		bank.NewAccountBalanceEntry(1, "address3", 1),
	}
	require.Equal(t, expected, bank.MergeAccountBalanceEntries(oldEntries, newEntries))
}

func TestMergeTokenAccountBalanceEntries(t *testing.T) {
	oldEntries := []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(0, "address1", 1),
		bank.NewTokenAccountBalanceEntry(0, "address2", 1),
	}
	newEntries := []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(1, "address1", 2),
		bank.NewTokenAccountBalanceEntry(1, "address3", 1),
	}
	expected := []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(1, "address1", 2),
		bank.NewTokenAccountBalanceEntry(0, "address2", 1),
		bank.NewTokenAccountBalanceEntry(1, "address3", 1),
	}
	require.Equal(t, expected, bank.MergeTokenAccountBalanceEntries(oldEntries, newEntries))
}
