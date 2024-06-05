package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"connectrpc.com/connect"
	"github.com/common-fate/clio/clierr"
	"github.com/common-fate/sdk/tokenstore"

	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	CurrentContext string `toml:"current_context" json:"current_context"`
	// Contexts allows multiple Common Fate tenancies to be switched between easily.
	// We don't have official support for this yet in the CLI,
	// but the config structure supports it so that it can be easily added in future.
	Contexts map[string]Context `toml:"context" json:"context"`
}

type Key string

var (
	APIURLKey           Key = "api_url"
	AccessURLKey        Key = "access_url"
	AuthzURLKey         Key = "authz_url"
	OIDCIssuerKey       Key = "oidc_issuer"
	OIDCClientIDKey     Key = "oidc_client_id"
	OIDCClientSecretKey Key = "oidc_client_secret"
	NameKey             Key = "name"
)

type Context struct {
	// the name of the context is the key of the toml entry
	name         string
	APIURL       string `toml:"api_url,omitempty" json:"api_url,omitempty"`
	AccessURL    string `toml:"access_url,omitempty" json:"access_url,omitempty"`
	AuthzURL     string `toml:"authz_url,omitempty" json:"authz_url,omitempty"`
	OIDCIssuer   string `toml:"oidc_issuer,omitempty" json:"oidc_issuer,omitempty"`
	OIDCClientID string `toml:"oidc_client_id,omitempty" json:"oidc_client_id,omitempty"`
	// OIDCClientSecret, if specified, will cause the client to use machine-to-machine OIDC authentication.
	OIDCClientSecret string `toml:"oidc_client_secret,omitempty" json:"oidc_client_secret,omitempty"`
	// WHEN ADDING ANY NEW CONFIG VARIABLES BELOW YOU NEED TO UPDATE THE FOLLOWING:
	// - Config Sources: config/file_source.go, config/env_source.go.
	// - New() function: config/load.go

	// HTTPClient is filled in by calling Initialize()
	HTTPClient connect.HTTPClient
	// HTTPClientBuilder is filled in by calling Initialize() and can be used when the transport needs to be modified
	HTTPClientBuilder func(transport http.RoundTripper) connect.HTTPClient

	// OIDCProvider is filled in by calling Initialize()
	OIDCProvider rp.RelyingParty

	TokenSource oauth2.TokenSource

	TokenStore TokenStore
}

// HTTPClient is the interface connect expects HTTP clients to implement. The
// standard library's *http.Client implements HTTPClient.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type TokenStore interface {
	Clear() error
	Save(token *oauth2.Token) error
	Token() (*oauth2.Token, error)
}

type InitializeOpts struct {
	TokenStore        TokenStore
	HttpClientWrapper func(c Doer, t TokenStore) Doer
}

func (c *Context) Initialize(ctx context.Context, opts InitializeOpts) error {
	// the OIDC client ID and issuer are *always* required, so make sure they've
	if c.OIDCClientID == "" {
		return errors.New("OIDC Client ID must be set. You can configure this by setting the CF_OIDC_CLIENT_ID environment variable, or specifying 'oidc_client_id' in the Common Fate config file (~/.cf/config by default). If you're using the Common Fate CLI, use 'cf configure' to set your config file up")
	}
	if c.OIDCIssuer == "" {
		return errors.New("OIDC Issuer must be set. You can configure this by setting the CF_OIDC_ISSUER environment variable, or specifying 'oidc_client_issuer' in the Common Fate config file (~/.cf/config by default). If you're using the Common Fate CLI, use 'cf configure' to set your config file up")
	}

	emptyClientSecret := ""
	scopes := []string{"openid", "email"}
	redirectURI := "http://localhost:18900/auth/callback"
	options := []rp.Option{
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}

	p, err := rp.NewRelyingPartyOIDC(c.OIDCIssuer, c.OIDCClientID, emptyClientSecret, redirectURI, scopes, options...)
	if err != nil {
		return err
	}

	// save the token in the keychain with the key OIDC Issuer + client ID.
	issuerURL, err := url.Parse(c.OIDCIssuer)
	if err != nil {
		return fmt.Errorf("invalid OIDC issuer url: %w", err)
	}
	issuerURL = issuerURL.JoinPath(c.OIDCClientID)

	c.OIDCProvider = p
	c.TokenStore = opts.TokenStore
	if c.TokenStore == nil && runtime.GOOS != "windows" {
		ts := tokenstore.New(tokenstore.Opts{
			Name: issuerURL.String(),
		})
		c.TokenStore = &ts
	}
	if c.TokenStore == nil && runtime.GOOS == "windows" {
		ts := tokenstore.NewWindows(tokenstore.Opts{
			Name: issuerURL.String(),
		})

		c.TokenStore = &ts
	}

	oauthconf := p.OAuthConfig()

	if c.OIDCClientSecret != "" {
		cfg := clientcredentials.Config{
			ClientID:     c.OIDCClientID,
			ClientSecret: c.OIDCClientSecret,
			TokenURL:     p.OAuthConfig().Endpoint.TokenURL,
		}

		_, err := cfg.Token(ctx)
		if err != nil {
			return err
		}
		c.TokenSource = cfg.TokenSource(ctx)

		c.HTTPClient = cfg.Client(ctx)
		if opts.HttpClientWrapper != nil {
			c.HTTPClient = opts.HttpClientWrapper(c.HTTPClient, c.TokenStore)
		}
		c.HTTPClientBuilder = func(transport http.RoundTripper) connect.HTTPClient {
			client := cfg.Client(ctx)
			client.Transport = transport
			if opts.HttpClientWrapper != nil {
				return opts.HttpClientWrapper(client, c.TokenStore)
			}
			return client
		}
		return nil
	}

	tok, err := c.TokenStore.Token()
	if err != nil {
		return err
	}

	src := &tokenstore.NotifyRefreshTokenSource{
		New:       oauthconf.TokenSource(ctx, tok),
		T:         tok,
		SaveToken: c.TokenStore.Save,
	}
	c.TokenSource = src

	c.HTTPClient = oauth2.NewClient(ctx, src)
	if opts.HttpClientWrapper != nil {
		c.HTTPClient = opts.HttpClientWrapper(c.HTTPClient, c.TokenStore)
	}

	c.HTTPClientBuilder = func(transport http.RoundTripper) connect.HTTPClient {
		client := oauth2.NewClient(ctx, src)
		client.Transport = transport
		if opts.HttpClientWrapper != nil {
			return opts.HttpClientWrapper(client, c.TokenStore)
		}
		return client
	}
	return nil
}

// Keys are all of the allowed keys in the Context section.
var Keys = []string{"oidc_issuer", "oidc_client_id", "api_url", "authz_url", "access_url", "oidc_client_secret"}

// Current loads the current context as specified in the 'current_context' field in the config file.
// It returns an error if there are no contexts, or if the 'current_context' field doesn't match
// any contexts defined in the config file.
func (c Config) Current() (*Context, error) {
	if c.Contexts == nil {
		return nil, clierr.New("No contexts were found in Common Fate config file.")
	}

	got, ok := c.Contexts[c.CurrentContext]
	if !ok {
		return nil, clierr.New(fmt.Sprintf("Could not find context '%s' in Common Fate config file", c.CurrentContext))
	}

	return &got, nil
}

// Default returns an empty config.
func Default() *Config {
	return &Config{
		CurrentContext: "default",
		Contexts: map[string]Context{
			"default": {},
		},
	}
}
