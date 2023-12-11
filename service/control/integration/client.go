package integration

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1/integrationv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) integrationv1alpha1connect.IntegrationServiceClient {
	return integrationv1alpha1connect.NewIntegrationServiceClient(cfg.HTTPClient, cfg.APIURL)
}
