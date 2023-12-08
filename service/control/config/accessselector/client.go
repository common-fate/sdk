package accessselector

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1/configv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) configv1alpha1connect.AccessSelectorServiceClient {
	return configv1alpha1connect.NewAccessSelectorServiceClient(cfg.HTTPClient, cfg.APIURL)
}
