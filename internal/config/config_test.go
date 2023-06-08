package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	// Set environment variables for testing

	os.Setenv("DRY_RUN", "true")
	os.Setenv("GITLAB_BASE_URL", "https://gitlab.com")
	os.Setenv("GITLAB_API_TOKEN", "token")
	os.Setenv("GITLAB_RUNNER_STATUS", "active")

	// Test that Load returns a Config struct with the expected values
	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "https://gitlab.com", cfg.BaseURL)
	assert.Equal(t, "token", cfg.Token)
	assert.Equal(t, "active", cfg.RunnerStatus)
	assert.True(t, cfg.DryRun)

	// Unset environment variables after testing
	os.Unsetenv("ENV_SECRET_PATH")
	os.Unsetenv("DRY_RUN")
	os.Unsetenv("GITLAB_BASE_URL")
	os.Unsetenv("GITLAB_API_TOKEN")
	os.Unsetenv("GITLAB_RUNNER_STATUS")

	// Remove the test .env file
	// os.Remove(".env.test")
}
