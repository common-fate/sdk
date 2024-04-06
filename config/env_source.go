package config

import (
	"os"
	"strings"
)

type EnvSource struct{}

// Load config variables.
func (ev EnvSource) Load(key Key) (string, error) {
	// turn the key into the environment variable,
	// e.g. "api_url" becomes "CF_API_URL"
	envVar := "CF_" + strings.ToUpper(string(key))

	return os.Getenv(envVar), nil
}
