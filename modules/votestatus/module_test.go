package votestatus_test

func (suite *ModuleTestSuite) TestModule_Name() {
	suite.Require().Equal("vote-status", suite.module.Name())
}
