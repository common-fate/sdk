package usage

import (
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/usage/v1alpha1/usagev1alpha1connect"
)

func NewFromConfig(cfg *factoryconfig.Context) usagev1alpha1connect.UsageServiceClient {
	return usagev1alpha1connect.NewUsageServiceClient(cfg.HTTPClient, cfg.BaseURL)
}
