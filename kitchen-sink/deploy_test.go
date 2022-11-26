package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func (suite *DeployTestSuite) TestLocalStackRunning() {
	suite.True(suite.isLocalStackRunning())
}

func (suite *DeployTestSuite) TestRunnerUrlExists() {
	suite.Log("Testing runner URL %s", suite.url)
	suite.NotEmpty(suite.url, "Runner URL is empty")
	_, err := url.ParseRequestURI(suite.url)
	suite.Nil(err, "Runner URL is not a valid URL")
}
