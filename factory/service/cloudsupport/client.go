package cloudsupport

import (
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/cloudsupport/v1alpha1/cloudsupportv1alpha1connect"
)

func NewFromConfig(cfg *factoryconfig.Context) cloudsupportv1alpha1connect.CloudSupportServiceClient {
	return cloudsupportv1alpha1connect.NewCloudSupportServiceClient(cfg.HTTPClient, cfg.BaseURL)
}
