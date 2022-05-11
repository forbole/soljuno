package votestatus_test

import (
	"fmt"

	"github.com/forbole/soljuno/modules/votestatus"
)

func (suite *ModuleTestSuite) TestModule_UpdateValidatorsStatus() {
	err := votestatus.UpdateValidatorsStatus(suite.db, suite.client)
	suite.Require().NoError(err)

	// with error client returns error
	errClient := suite.client.GetCached()
	errClient.WithError(fmt.Errorf("error"))
	err = votestatus.UpdateValidatorsStatus(suite.db, &errClient)
	suite.Require().Error(err)

	// with error db returns error
	errDb := suite.db.GetCached()
	errDb.WithError(fmt.Errorf("error"))
	err = votestatus.UpdateValidatorsStatus(&errDb, suite.client)
	suite.Require().Error(err)
}

func (suite *ModuleTestSuite) TestModule_UpdateValidatorSkipRates() {
	err := votestatus.UpdateValidatorSkipRates(1, suite.db, suite.client)
	suite.Require().NoError(err)

	// with error client returns error
	errClient := suite.client.GetCached()
	errClient.WithError(fmt.Errorf("error"))
	err = votestatus.UpdateValidatorSkipRates(1, suite.db, &errClient)
	suite.Require().Error(err)

	// with error db returns error
	errDb := suite.db.GetCached()
	errDb.WithError(fmt.Errorf("error"))
	err = votestatus.UpdateValidatorSkipRates(1, &errDb, suite.client)
	suite.Require().Error(err)
}

func (suite *ModuleTestSuite) TestModule_GetSkipRateReference() {
	m := make(map[int]bool)
	m[0] = true
	total, skip := votestatus.GetSkipRateReference(1, m, []int{0, 1})
	suite.Require().Equal(2, total)
	suite.Require().Equal(1, skip)
}
