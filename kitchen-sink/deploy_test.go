package test

func (suite *DeployTestSuite) TestLocalStackRunning() {
	suite.True(suite.isLocalStackRunning())
}
