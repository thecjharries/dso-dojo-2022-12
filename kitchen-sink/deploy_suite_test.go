package test

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type LocalStackStatus struct {
	Running bool `json:"running"`
}

type DeployTestSuite struct {
	suite.Suite
	options *terraform.Options
	logger  *logger.Logger
}

func (suite *DeployTestSuite) Log(message string, args ...interface{}) {
	suite.logger.Logf(suite.T(), message, args...)
}

func (suite *DeployTestSuite) isLocalStackRunning() bool {
	command := exec.Command("localstack", "status", "docker", "--format", "json")
	stdout, err := command.Output()
	suite.Nil(err, "Error running localstack status command")
	var status LocalStackStatus
	err = json.Unmarshal(stdout, &status)
	suite.Nil(err, "Error unmarshalling localstack status")
	return status.Running
}

func (suite *DeployTestSuite) ensureLocalStackRunning() {
	if !suite.isLocalStackRunning() {
		suite.Log("LocalStack is not running, starting it now...")
		command := exec.Command("localstack", "start", "--detached")
		err := command.Run()
		suite.Nil(err, "Error running localstack start command")
	}
	suite.Log("LocalStack is running")
}

func (suite *DeployTestSuite) ensureLocalStackStopped() {
	if suite.isLocalStackRunning() {
		suite.Log("LocalStack is running, stopping it now...")
		command := exec.Command("localstack", "stop")
		err := command.Run()
		suite.Nil(err, "Error running localstack stop command")
	}
	suite.Log("LocalStack is stopped")
}

func (suite *DeployTestSuite) SetupSuite() {
	suite.logger = logger.Terratest
	tmpTestFolder := test_structure.CopyTerraformFolderToTemp(suite.T(), ".", "module")
	suite.options = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformBinary: "tflocal",
		TerraformDir:    tmpTestFolder,
	})
	suite.ensureLocalStackRunning()
}

func (suite *DeployTestSuite) TearDownSuite() {
	suite.ensureLocalStackStopped()
}

func TestDeploySuite(t *testing.T) {
	suite.Run(t, new(DeployTestSuite))
}
