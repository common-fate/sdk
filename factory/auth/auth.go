package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/common-fate/cloudproof/aws/stsproof"
	"golang.org/x/oauth2"
)

// deploymentTransport adds additional headers used by the Factory OIDC server
// to the outgoing HTTP requests.
type deploymentTransport struct {
	LicenceKey     string
	DeploymentName string
}

func (d *deploymentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if d.LicenceKey != "" {
		r.Header.Add("X-Common-Fate-Licence-Key", d.LicenceKey)
	}

	if d.DeploymentName != "" {
		r.Header.Add("X-Common-Fate-Deployment-Name", d.DeploymentName)
	}

	return http.DefaultTransport.RoundTrip(r)
}

type AuthClient struct {
	client *http.Client
	issuer *url.URL
	config *Config
	ctx    context.Context
}

type Opts struct {
	LicenceKey     string
	DeploymentName string
	Issuer         string
	BaseClient     *http.Client
}

// NewClient creates a new authentication client.
// The context on the client should be long-lived, e.g. context.Background().
func NewClient(ctx context.Context, opts Opts) (*AuthClient, error) {
	if opts.BaseClient == nil {
		opts.BaseClient = &http.Client{
			Transport: &deploymentTransport{
				LicenceKey:     opts.LicenceKey,
				DeploymentName: opts.DeploymentName,
			},
		}
	}

	issuer, err := url.Parse(opts.Issuer)
	if err != nil {
		return nil, fmt.Errorf("invalid issuer url '%s': %w", opts.Issuer, err)
	}

	c := AuthClient{
		client: opts.BaseClient,
		issuer: issuer,
		ctx:    ctx,
	}

	c.config, err = c.Discover(ctx)
	if err != nil {
		return nil, fmt.Errorf("error calling OIDC discovery endpoint: %w", err)
	}

	return &c, nil
}

// Token retrieves a new oauth2 Token from the endpoint.
//
// It is compatible with the TokenSource interface from the x/oauth2 library.
func (c *AuthClient) Token() (*oauth2.Token, error) {
	proof, err := stsproof.New(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating AWS identity proof: %w", err)
	}

	form := url.Values{}
	form.Add("grant_type", "https://commonfate.io/cloudproof/aws")
	form.Add("timestamp", proof.Time.Format(time.RFC3339))
	form.Add("security_token", proof.SecurityToken)
	form.Add("signature", proof.Signature)

	req, err := http.NewRequest("POST", c.config.TokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling token endpoint: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 300 {
		bodyString, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("received status code %v from token endpoint and could not parse body: %w", res.StatusCode, err)
		}

		return nil, fmt.Errorf("token endpoint error: %s", string(bodyString))
	}

	var token oauth2.Token

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling token: %w", err)
	}

	return &token, nil
}

func (c *AuthClient) Config() *Config {
	return c.config
}

type Config struct {
	Issuer                 string   `json:"issuer"`
	TokenEndpoint          string   `json:"token_endpoint"`
	GrantTypesSupported    []string `json:"grant_types_supported"`
	ResponseTypesSupported []string `json:"response_types_supported"`
}

func (c *AuthClient) Discover(ctx context.Context) (*Config, error) {
	configURL := c.issuer.JoinPath(".well-known", "openid-configuration")

	req, err := http.NewRequestWithContext(ctx, "GET", configURL.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cfg Config

	err = json.NewDecoder(res.Body).Decode(&cfg)
	return &cfg, err
}
