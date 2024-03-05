package reset

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/integration/reset/v1alpha1/resetv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) resetv1alpha1connect.ResetServiceClient {
	return resetv1alpha1connect.NewResetServiceClient(cfg.HTTPClient, cfg.APIURL)
}
