package bank_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules/bank"
)

type ModuleTestSuite struct {
	suite.Suite
	module *bank.Module
}

type MockDb struct{}

var _ db.BankDb = &MockDb{}

func (db MockDb) SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error {
	return nil
}
func (db MockDb) SaveAccountTokenBalances(slot uint64, accounts []string, balances []uint64) error {
	return nil
}
func (db MockDb) SaveAccountHistoryBalances(time time.Time, accounts []string, balances []uint64) error {
	return nil
}
func (db MockDb) SaveAccountHistoryTokenBalances(time time.Time, accounts []string, balances []uint64) error {
	return nil
}

func TestModuleTestSuitee(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = bank.NewModule(MockDb{})
}
