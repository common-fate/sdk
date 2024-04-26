// Package factoryconfig holds configuration for the Factory service SDK.
// This service is hosted by Common Fate, so configuring a client is slightly
// different than configuring a client for a Common Fate deployment.
package factoryconfig

import (
	"context"
	"net/http"

	"github.com/common-fate/sdk/factory/auth"
	"golang.org/x/oauth2"
)

type Opts struct {
	LicenceKey     string
	DeploymentName string

	// BaseURL of the Factory service to connect to.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	BaseURL string

	// OIDCIssuer is the OIDC issuer to use.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	OIDCIssuer string
}

type Context struct {
	// BaseURL of the Factory service to connect to.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	BaseURL string

	// OIDCIssuer is the OIDC issuer to use.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	OIDCIssuer string

	// HTTPClient is filled in by calling Initialize()
	HTTPClient *http.Client
}

// Load and initialize a client context. You can override values by providing them in 'opts'.
func Load(ctx context.Context, opts Opts) (*Context, error) {
	if opts.BaseURL == "" {
		opts.BaseURL = "https://factory.commonfate.io"
	}

	if opts.OIDCIssuer == "" {
		opts.OIDCIssuer = "https://factory.commonfate.io"
	}

	c := Context{
		BaseURL:    opts.BaseURL,
		OIDCIssuer: opts.OIDCIssuer,
	}

	authClient, err := auth.NewClient(ctx, auth.Opts{
		Issuer: c.BaseURL,
	})
	if err != nil {
		return nil, err
	}

	c.HTTPClient = oauth2.NewClient(ctx, authClient)

	return &c, nil
}
