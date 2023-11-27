package config

import (
	"context"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/common-fate/clio"
)

// LoadDefault is a shorthand function which
// calls Load() to load the configuration,
// and then calls cfg.Current()
// to get the current API context.
// It returns an error if either method fails.
func LoadDefault(ctx context.Context) (*Context, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}

	current, err := cfg.Current()
	if err != nil {
		return nil, err
	}

	err = current.Initialize(ctx)
	if err != nil {
		return nil, err
	}

	return current, nil
}

type ServiceAccount struct {
	APIURL    string
	IssuerURL string

	ClientId     string
	ClientSecret string
}

func NewServiceAccount(ctx context.Context, ServiceAccount ServiceAccount) (*Context, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}

	current, err := cfg.Current()
	if err != nil {
		return nil, err
	}

	current.APIURL = ServiceAccount.APIURL
	current.AccessURL = ServiceAccount.IssuerURL
	current.ServiceAccountClientID = &ServiceAccount.ClientId
	current.ServiceAccountClientSecret = &ServiceAccount.ClientSecret

	err = current.Initialize(ctx)
	if err != nil {
		return nil, err
	}

	return current, nil

}

func load() (*Config, error) {
	// if COMMON_FATE_CIEM_CONFIG_FILE is set, use a custom file path
	// for the config file location.
	// the file specified must exist.
	customPath := os.Getenv("COMMON_FATE_CIEM_CONFIG_FILE")
	if customPath != "" {
		return openConfigFile(customPath)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fp := filepath.Join(home, ".commonfate", "ciem")
	cfg, err := openConfigFile(fp)
	if os.IsNotExist(err) {
		// return an empty config if the file doesn't exist
		return Default(), nil
	}
	if err != nil {
		// otherwise if we get an error, return it
		return nil, err
	}

	return cfg, nil
}

func openConfigFile(filepath string) (*Config, error) {
	clio.Debugw("loading config", "path", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config

	_, err = toml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	clio.Debugw("loaded config", "cfg", cfg)
	return &cfg, nil
}
