package config

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/common-fate/ciem/tokenstore"
	"github.com/common-fate/clio/clierr"
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

type Context struct {
	// the name of the context is the key of the toml entry
	name                       string
	APIURL                     string  `toml:"api_url,omitempty" json:"api_url,omitempty"`
	AccessURL                  string  `toml:"access_url,omitempty" json:"access_url,omitempty"`
	OIDCIssuer                 string  `toml:"oidc_issuer,omitempty" json:"oidc_issuer,omitempty"`
	OIDCClientID               string  `toml:"oidc_client_id,omitempty" json:"oidc_client_id,omitempty"`
	ServiceAccountClientID     *string `toml:"service_account_client_id,omitempty" json:"service_account_client_id,omitempty"`
	ServiceAccountClientSecret *string `toml:"service_account_client_secret,omitempty" json:"service_account_client_secret,omitempty"`

	// HTTPClient is filled in by calling Initialize()
	HTTPClient *http.Client

	// OIDCProvider is filled in by calling Initialize()
	OIDCProvider rp.RelyingParty

	TokenStore tokenstore.Storage
}

func (c *Context) Initialize(ctx context.Context) error {
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

	c.OIDCProvider = p
	c.TokenStore = tokenstore.New(tokenstore.Opts{
		Name: c.name,
	})

	oauthconf := p.OAuthConfig()

	if c.ServiceAccountClientID != nil {
		cfg := clientcredentials.Config{
			ClientID:     *c.ServiceAccountClientID,
			ClientSecret: *c.ServiceAccountClientSecret,
			Scopes:       []string{"cf.client/read", "cf.client/write"},
			TokenURL:     c.AccessURL + "/oauth2/token",
		}

		_, err := cfg.Token(ctx)
		if err != nil {
			return err
		}

		c.HTTPClient = cfg.Client(ctx)
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

	c.HTTPClient = oauth2.NewClient(ctx, src)

	return nil
}

// Keys are all of the allowed keys in the Context section.
var Keys = []string{"oidc_issuer", "oidc_client_id", "api_url"}

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
		CurrentContext: "",
		Contexts:       map[string]Context{},
	}
}
