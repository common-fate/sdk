package config

import (
	"os"
	"strings"
)

type EnvSource struct{}

// Load config variables.
// The function must not set config variables if they are already set.
func (ev EnvSource) Load(key Key) string {
	// turn the key into the environment variable,
	// e.g. "api_url" becomes "CF_API_URL"
	envVar := "CF_" + strings.ToUpper(string(key))

	return os.Getenv(envVar)
}
