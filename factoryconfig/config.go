// Package factoryconfig holds configuration for the Factory service SDK.
// This service is hosted by Common Fate, so configuring a client is slightly
// different than configuring a client for a Common Fate deployment.
package factoryconfig

import (
	"context"
	"net/http"
)

type Opts = Context // held to avoid breaking changes here, in future the structs may become different.

type Context struct {
	// BaseURL of the Factory service to connect to.
	// Defaults to "https://factory.commonfate.io"
	// if not provided.
	BaseURL string
	// A custom HTTP Client. Defaults to http.DefaultClient
	// if not provided.
	HTTPClient *http.Client
}

// LoadDefault loads a client context with all defaults configured.
func LoadDefault(ctx context.Context) *Context {
	return Load(ctx, Opts{})
}

// Load a client context. You can override values by providing them in 'opts'.
func Load(ctx context.Context, opts Opts) *Context {
	if opts.BaseURL == "" {
		opts.BaseURL = "https://factory.commonfate.io"
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}

	c := Context{
		BaseURL:    opts.BaseURL,
		HTTPClient: opts.HTTPClient,
	}

	return &c
}
