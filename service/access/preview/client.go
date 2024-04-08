package preview

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/access/v1alpha1/accessv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accessv1alpha1connect.PreviewServiceClient {
	return accessv1alpha1connect.NewPreviewServiceClient(cfg.HTTPClient, cfg.AccessURL)
}
