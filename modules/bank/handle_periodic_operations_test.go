package bank_test

import "github.com/forbole/soljuno/modules/bank"

func (suite *ModuleTestSuite) TestModule_HandlePeriodicOperations() {
	suite.module.BalanceEntries = []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(0, "address", 10),
	}
	suite.module.TokenBalanceEntries = []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(0, "address", 10),
	}

	suite.module.HandlePeriodicOperations()
	suite.Require().Empty(suite.module.BalanceEntries)
	suite.Require().Empty(suite.module.TokenBalanceEntries)
}
