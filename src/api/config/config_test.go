package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnv(t *testing.T) {
	readEnv()
	assert.Equal(t, "development", os.Getenv("APPENV"))
}

func TestLoadConfig(t *testing.T) {
	LoadConfig()
	assert.Equal(t, "abc123", Config.GithubAccessKey)
	assert.Equal(t, "localhost:3000", Config.AppUrl)
}
