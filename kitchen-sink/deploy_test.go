package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gruntwork-io/terratest/modules/retry"
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

type RunnerResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Link    string `json:"link"`
	Message string `json:"message"`
}

func (suite *DeployTestSuite) TestRunnerValidUnused() {
	_, err := retry.DoWithRetryE(suite.T(), "Testing code ABCDE-12345-FGIJH", 6, 10*time.Second, func() (string, error) {
		var requestBody = []byte(`{"code":"ABCDE-12345-FGIJH"}`)
		response, responseErr := http.Post(suite.url, "application/json", bytes.NewBuffer(requestBody))
		if responseErr != nil {
			return "", responseErr
		}
		defer response.Body.Close()
		var runnerResponse RunnerResponse
		decodeErr := json.NewDecoder(response.Body).Decode(&runnerResponse)
		suite.Nil(decodeErr, "Error decoding response")
		if strings.HasPrefix(runnerResponse.Message, "exception") {
			return "", errors.New(runnerResponse.Message)
		}
		suite.Log("Runner response: %+v", runnerResponse)
		suite.True(runnerResponse.Success, "Runner response was not successful")
		return "", nil
	})
	suite.Nil(err, "Error making request to runner")
}
