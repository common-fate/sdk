package configsvc

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1/configv1alpha1connect"
)

type Client struct {
	cfg *config.Context
}

func NewFromConfig(cfg *config.Context) *Client {
	return &Client{cfg: cfg}
}

func (c Client) Selector() configv1alpha1connect.SelectorServiceClient {
	return configv1alpha1connect.NewSelectorServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}

func (c Client) AccessWorkflow() configv1alpha1connect.AccessWorkflowServiceClient {
	return configv1alpha1connect.NewAccessWorkflowServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}

func (c Client) AvailabilitySpec() configv1alpha1connect.AvailabilitySpecServiceClient {
	return configv1alpha1connect.NewAvailabilitySpecServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}

func (c Client) WebhookProvisioner() configv1alpha1connect.WebhookProvisionerServiceClient {
	return configv1alpha1connect.NewWebhookProvisionerServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}
