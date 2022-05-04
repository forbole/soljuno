package bank_test

import (
	"fmt"

	"github.com/forbole/soljuno/modules/bank"
)

func (suite *ModuleTestSuite) TestModule_HandlePeriodicOperations() {
	suite.module.BalanceEntries = []bank.AccountBalanceEntry{
		bank.NewAccountBalanceEntry(0, "address", 10),
	}
	suite.module.TokenBalanceEntries = []bank.TokenAccountBalanceEntry{
		bank.NewTokenAccountBalanceEntry(0, "address", 10),
	}
	brokenDb := suite.db.GetCached()
	brokenDb.WithError(fmt.Errorf("error"))
	err := bank.HandlePeriodicOperations(suite.module, brokenDb)
	suite.Require().Error(err)
	suite.Require().NotEmpty(suite.module.BalanceEntries)
	suite.Require().NotEmpty(suite.module.TokenBalanceEntries)

	err = bank.HandlePeriodicOperations(suite.module, suite.db)
	suite.Require().NoError(err)
	suite.Require().Empty(suite.module.BalanceEntries)
	suite.Require().Empty(suite.module.TokenBalanceEntries)
}
