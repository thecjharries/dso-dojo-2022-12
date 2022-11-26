package test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
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
	setupSuccess bool
	logger       *logger.Logger
	options      *terraform.Options
	plan         *terraform.PlanStruct
	url          string
}

func (suite *DeployTestSuite) Log(message string, args ...interface{}) {
	suite.logger.Logf(suite.T(), message, args...)
}

func (suite *DeployTestSuite) isLocalStackRunning() bool {
	command := exec.Command("localstack", "status", "docker", "--format", "json")
	command.Env = os.Environ()
	command.Env = append(command.Env, "TF_COMPAT_MODE=true")
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
	// https://github.com/stretchr/testify/issues/1123#issuecomment-964729053
	defer (func() {
		if !suite.setupSuccess {
			suite.TearDownSuite()
		}
	})()
	suite.ensureLocalStackRunning()
	suite.logger = logger.Terratest
	tmpTestFolder := test_structure.CopyTerraformFolderToTemp(suite.T(), ".", "module")
	planFilePath := filepath.Join(tmpTestFolder, "..", "plan.out")
	suite.options = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		TerraformBinary: "tflocal",
		TerraformDir:    tmpTestFolder,
		PlanFilePath:    planFilePath,
	})
	var planErr error
	suite.plan, planErr = terraform.InitAndPlanAndShowWithStructE(suite.T(), suite.options)
	suite.Nil(planErr, "Error running terraform init and plan")
	applyOutput, applyErr := terraform.ApplyAndIdempotentE(suite.T(), suite.options)
	suite.Nil(applyErr, "Error running terraform apply")
	suite.Log("Apply output: %s", applyOutput)
	suite.url = terraform.Output(suite.T(), suite.options, "runner_url")
	suite.setupSuccess = true
}

func (suite *DeployTestSuite) TearDownSuite() {
	terraform.Destroy(suite.T(), suite.options)
	suite.ensureLocalStackStopped()
}

func TestDeploySuite(t *testing.T) {
	suite.Run(t, new(DeployTestSuite))
}
