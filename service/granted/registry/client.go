package registry

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/granted/registry/aws/v1alpha1/awsv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) awsv1alpha1connect.ProfileRegistryServiceClient {
	return awsv1alpha1connect.NewProfileRegistryServiceClient(cfg.HTTPClient, cfg.APIURL)
}
