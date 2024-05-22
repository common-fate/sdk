package deployment

import (
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/deployment/v1alpha1/deploymentv1alpha1connect"
)

func NewFromConfig(cfg *factoryconfig.Context) deploymentv1alpha1connect.DeploymentServiceClient {
	return deploymentv1alpha1connect.NewDeploymentServiceClient(cfg.HTTPClient, cfg.BaseURL)
}
