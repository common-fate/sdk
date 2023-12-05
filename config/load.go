package config

import (
	"context"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/common-fate/clio"
	"github.com/common-fate/sdk/tokenstore"
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

	err = current.Initialize(ctx, InitializeOpts{})
	if err != nil {
		return nil, err
	}

	return current, nil
}

type Opts struct {
	APIURL       string
	AccessURL    string
	ClientID     string
	ClientSecret string
	OIDCIssuer   string
}

func New(ctx context.Context, opts Opts) (*Context, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}

	current, err := cfg.Current()
	if err != nil {
		return nil, err
	}

	current.APIURL = opts.APIURL
	current.AccessURL = opts.AccessURL
	current.OIDCClientID = opts.ClientID
	current.OIDCClientSecret = &opts.ClientSecret
	current.OIDCIssuer = opts.OIDCIssuer

	err = current.Initialize(ctx, InitializeOpts{})
	if err != nil {
		return nil, err
	}

	return current, nil
}

// NewServerContext requires all option to be passed and does not attempt to read from the local config file
// it also uses an in memory token store to avoid keychain access
// usefull for service to service client credentials auth
// targetServiceURL shoudl be the service you will be calling with these credentials
// if you need to call multiple then make multiple clients
// @TODO improve this
func NewServerContext(ctx context.Context, targetServiceURL string, opts Opts) (*Context, error) {

	context := &Context{
		APIURL:           opts.APIURL,
		AccessURL:        opts.AccessURL,
		OIDCClientID:     opts.ClientID,
		OIDCClientSecret: &opts.ClientSecret,
		OIDCIssuer:       opts.OIDCIssuer,
	}

	// Initialise with an in memory token store to avoid keychain use
	err := context.Initialize(ctx, InitializeOpts{TokenStore: tokenstore.NewInMemoryTokenStore(), ServiceURL: targetServiceURL})
	if err != nil {
		return nil, err
	}

	return context, nil
}

func load() (*Config, error) {
	// if COMMON_FATE_CONFIG_FILE is set, use a custom file path
	// for the config file location.
	// the file specified must exist.
	customPath := os.Getenv("COMMON_FATE_CONFIG_FILE")
	if customPath != "" {
		return openConfigFile(customPath)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fp := filepath.Join(home, ".commonfate", "config")
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
