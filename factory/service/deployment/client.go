package deployment

import (
	"connectrpc.com/connect"
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/deployment/v1alpha1/deploymentv1alpha1connect"
)

func NewFromConfig(cfg *factoryconfig.Context, opts ...connect.ClientOption) deploymentv1alpha1connect.DeploymentServiceClient {
	return deploymentv1alpha1connect.NewDeploymentServiceClient(cfg.HTTPClient, cfg.BaseURL, opts...)
}
