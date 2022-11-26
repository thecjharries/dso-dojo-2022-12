package test

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

type LocalStackStatus struct {
	Running bool `json:"running"`
}

func isLocalStackRunning(t *testing.T) bool {
	command := exec.Command("localstack", "status", "docker", "--format", "json")
	stdout, err := command.Output()
	assert.Nil(t, err, "Error running localstack status command")
	var status LocalStackStatus
	err = json.Unmarshal(stdout, &status)
	assert.Nil(t, err, "Error unmarshalling localstack status")
	return status.Running
}

func ensureLocalStack(t *testing.T) {
	if !isLocalStackRunning(t) {
		logger.Log(t, "LocalStack is not running, starting it now...")
		command := exec.Command("localstack", "start", "--detached")
		err := command.Run()
		assert.Nil(t, err, "Error running localstack start command")
	}
	logger.Log(t, "LocalStack is running")
}

func TestLocalStackRunning(t *testing.T) {
	ensureLocalStack(t)
	assert.True(t, isLocalStackRunning(t))
}
