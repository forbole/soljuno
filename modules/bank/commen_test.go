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
	db     *MockDb
}

type MockDb struct {
	err error
}

var _ db.BankDb = &MockDb{}

func (db MockDb) SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error {
	return db.err
}
func (db MockDb) SaveAccountTokenBalances(slot uint64, accounts []string, balances []uint64) error {
	return db.err
}
func (db MockDb) SaveAccountHistoryBalances(time time.Time, accounts []string, balances []uint64) error {
	return db.err
}
func (db MockDb) SaveAccountHistoryTokenBalances(time time.Time, accounts []string, balances []uint64) error {
	return db.err
}

func (m MockDb) GetCached() MockDb {
	return m
}

func (m *MockDb) WithError(err error) {
	m.err = err
}

func TestModuleTestSuitee(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = bank.NewModule(&MockDb{})
	suite.db = &MockDb{}
}
