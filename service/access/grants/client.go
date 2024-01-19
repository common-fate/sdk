package request

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/access/v1alpha1/accessv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accessv1alpha1connect.GrantsServiceClient {
	return accessv1alpha1connect.NewGrantsServiceClient(cfg.HTTPClient, cfg.AccessURL)
}
