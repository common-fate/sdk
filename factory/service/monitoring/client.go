package monitoring

import (
	"connectrpc.com/connect"
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/monitoring/v1alpha1/monitoringv1alpha1connect"
)

type Client struct {
	cfg  *factoryconfig.Context
	opts []connect.ClientOption
}

func NewFromConfig(cfg *factoryconfig.Context, opts ...connect.ClientOption) *Client {
	return &Client{cfg: cfg, opts: opts}
}

func (c *Client) Tokens() monitoringv1alpha1connect.TokenServiceClient {
	return monitoringv1alpha1connect.NewTokenServiceClient(c.cfg.HTTPClient, c.cfg.BaseURL, c.opts...)
}

func (c *Client) Validation() monitoringv1alpha1connect.ValidationServiceClient {
	return monitoringv1alpha1connect.NewValidationServiceClient(c.cfg.HTTPClient, c.cfg.BaseURL, c.opts...)
}
