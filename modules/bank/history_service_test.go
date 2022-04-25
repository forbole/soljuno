package bank_test

import "github.com/forbole/soljuno/modules/bank"

func (suite *ModuleTestSuite) TestModule_ExecHistory() {
	suite.module.HistoryBalanceEntries = []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(0, "address", 10),
	}
	suite.module.HistoryTokenBalanceEntries = []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(0, "address", 10),
	}

	err := suite.module.ExecHistory()
	suite.Require().NoError(err)
	suite.Require().Empty(suite.module.HistoryBalanceEntries)
	suite.Require().Empty(suite.module.HistoryTokenBalanceEntries)

}
