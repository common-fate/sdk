// Package factoryconfig holds configuration for the Factory service SDK.
// This service is hosted by Common Fate, so configuring a client is slightly
// different than configuring a client for a Common Fate deployment.
package factoryconfig

import (
	"context"
	"net/http"

	"github.com/common-fate/sdk/factory/auth"
)

type Opts struct {
	LicenceKey string

	// BaseURL of the Factory service to connect to.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	BaseURL string
}

type Context struct {
	// BaseURL of the Factory service to connect to.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	BaseURL string

	// HTTPClient is filled in by calling Load()
	HTTPClient *http.Client
}

// Load and initialize a client context. You can override values by providing them in 'opts'.
func Load(ctx context.Context, opts Opts) (*Context, error) {
	if opts.BaseURL == "" {
		opts.BaseURL = "https://factory.commonfate.io"
	}

	c := Context{
		BaseURL: opts.BaseURL,
	}

	c.HTTPClient = &http.Client{
		Transport: &auth.Transport{
			LicenceKey: opts.LicenceKey,
		},
	}

	return &c, nil
}
