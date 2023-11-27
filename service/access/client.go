package access

import (
	"github.com/common-fate/ciem/config"
	"github.com/common-fate/sdk/gen/commonfate/access/v1alpha1/accessv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accessv1alpha1connect.AccessServiceClient {
	return accessv1alpha1connect.NewAccessServiceClient(cfg.HTTPClient, cfg.AccessURL)
}
