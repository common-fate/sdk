package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/common-fate/clio"
	"github.com/common-fate/sdk/tokenstore"
)

// LoadDefault is a shorthand function which loads the default SDK configuration by calling New()
// with default values.
//
// By default, config is loaded from environment variables, and then falling back to a TOML configuration
// file for any config values which are not set.
//
// The file is ~/.cf/config by default, but this can be overridden with the CF_CONFIG_PATH environment variable.
//
// Environment variables follow the pattern 'CF_CONFIG_VARIABLE_NAME', where CONFIG_VARIABLE_NAME
// is the name of the configuration value in upper snake-case.
// For example, 'CF_API_URL'.
func LoadDefault(ctx context.Context) (*Context, error) {
	return New(ctx, Opts{})
}

type Opts struct {
	APIURL string
	// AccessURL is the base URL of the Access Handler.
	// If empty, the API URL is used.
	AccessURL string
	// AccessURL is the base URL of the authz service.
	// If empty, the API URL is used.
	AuthzURL     string
	ClientID     string
	ClientSecret string
	OIDCIssuer   string

	// The token storage backend to use for OIDC tokens.
	// If not provided, will use the keychain backend.
	TokenStore TokenStore

	// the config sources to load config from.
	// Must be either 'env' or 'file'.
	// Defaults to ['env', 'file'] if not provided.
	ConfigSources []string
}

type configSource interface {
	Load(key Key) string
}

func loadFromSources(value *string, key Key, sources []configSource) {
	if value == nil {
		return
	}

	if *value != "" {
		return
	}

	for _, source := range sources {
		loadedValue := source.Load(key)
		if loadedValue != "" {
			*value = loadedValue
			return
		}
	}
}

// New creates an initialised SDK context to be used for creating SDK clients.
// For example:
//
//	import "github.com/common-fate/sdk/config"
//	import "github.com/common-fate/sdk/service/access"
//
//	cfg, err := config.New(ctx, opts{})
//	if err != nil {
//		return err
//	}
//	client := access.NewClient(cfg)
//
// Configuration values, such as the API URL and OIDC Client ID to use,
// can be provided via the Opts{} argument.
//
// The New() method will look configuration values up from environment variables
// and the config file (~/.cf/config by default) if they are not provided in Opts{}.
// By default, the order of priority is:
//
// 1. Static value, set in Opts{}
//
// 2. Environment variable
//
// 3. Config file
//
// This behaviour can be customised by setting opts.ConfigSources.
//
// Environment variables follow the pattern 'CF_CONFIG_VARIABLE_NAME', where CONFIG_VARIABLE_NAME
// is the name of the configuration value in upper snake-case.
// For example, 'CF_API_URL'.
func New(ctx context.Context, opts Opts) (*Context, error) {
	// set up a default config
	cfg := Context{
		APIURL:           opts.APIURL,
		AccessURL:        opts.AccessURL,
		AuthzURL:         opts.AuthzURL,
		OIDCClientID:     opts.ClientID,
		OIDCIssuer:       opts.OIDCIssuer,
		OIDCClientSecret: opts.ClientSecret,
	}

	if opts.ConfigSources == nil {
		opts.ConfigSources = []string{"env", "file"}
	}

	var sources []configSource

	for _, loaderType := range opts.ConfigSources {
		switch loaderType {
		case "env":
			sources = append(sources, EnvSource{})
		case "file":
			sources = append(sources, &FileSource{})

		default:
			return nil, fmt.Errorf("invalid config loader: %s (valid types are 'env' and 'file')", loaderType)
		}
	}

	loadFromSources(&cfg.APIURL, APIURLKey, sources)
	loadFromSources(&cfg.AuthzURL, AuthzURLKey, sources)
	loadFromSources(&cfg.AccessURL, AccessURLKey, sources)
	loadFromSources(&cfg.OIDCClientID, OIDCClientIDKey, sources)
	loadFromSources(&cfg.OIDCClientSecret, OIDCClientSecretKey, sources)
	loadFromSources(&cfg.OIDCIssuer, OIDCIssuerKey, sources)

	err := cfg.Initialize(ctx, InitializeOpts{TokenStore: opts.TokenStore})
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// NewServerContext requires all option to be passed and does not attempt to read from the local config file
// it also uses an in memory token store to avoid keychain access
func NewServerContext(ctx context.Context, opts Opts) (*Context, error) {
	if opts.ConfigSources == nil {
		// don't source from from file
		opts.ConfigSources = []string{"env"}
	}

	if opts.TokenStore == nil {
		// default to an in-memory token store
		opts.TokenStore = tokenstore.NewInMemoryTokenStore()
	}

	return New(ctx, opts)
}

func load() (*Config, error) {
	// if CF_CONFIG_FILE is set, use a custom file path
	// for the config file location.
	// the file specified must exist.
	customPath := os.Getenv("CF_CONFIG_FILE")
	if customPath != "" {
		return openConfigFile(customPath)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fp := filepath.Join(home, ".cf", "config")
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

func ListContexts() ([]string, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}
	contexts := []string{}
	for k := range cfg.Contexts {
		contexts = append(contexts, k)
	}
	return contexts, nil
}

func SwitchContext(contextName string) error {
	cfg, err := load()
	if err != nil {
		return err
	}
	if _, ok := cfg.Contexts[contextName]; ok {
		cfg.CurrentContext = contextName
		return Save(cfg)
	}
	return errors.New("context not found in config file")
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

	for k, context := range cfg.Contexts {
		context.name = k
		cfg.Contexts[k] = context
	}

	clio.Debugw("loaded config", "cfg", cfg)
	return &cfg, nil
}
