package config

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/common-fate/clio"
	"github.com/common-fate/grab"
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
}

// in order of precedence, select a source for the client ID or nil
func clientSecret(contextSecret *string, opts Opts) *string {
	clientSecret := grab.FirstNonZero(opts.ClientSecret, os.Getenv("CF_OIDC_CLIENT_SECRET"), grab.Value(contextSecret))
	if clientSecret == "" {
		return nil
	}
	return &clientSecret
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

	current.OIDCClientSecret = clientSecret(current.OIDCClientSecret, opts)

	current.OIDCIssuer = opts.OIDCIssuer

	err = current.Initialize(ctx, InitializeOpts{})
	if err != nil {
		return nil, err
	}

	return current, nil
}

// NewServerContext requires all option to be passed and does not attempt to read from the local config file
// it also uses an in memory token store to avoid keychain access
func NewServerContext(ctx context.Context, opts Opts) (*Context, error) {

	context := &Context{
		APIURL:       opts.APIURL,
		AccessURL:    opts.AccessURL,
		AuthzURL:     opts.AuthzURL,
		OIDCClientID: opts.ClientID,

		OIDCIssuer: opts.OIDCIssuer,
	}
	context.OIDCClientSecret = clientSecret(context.OIDCClientSecret, opts)

	// Initialise with an in memory token store to avoid keychain use
	err := context.Initialize(ctx, InitializeOpts{TokenStore: tokenstore.NewInMemoryTokenStore()})
	if err != nil {
		return nil, err
	}

	return context, nil
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

	clio.Debugw("loaded config", "cfg", cfg)
	return &cfg, nil
}
