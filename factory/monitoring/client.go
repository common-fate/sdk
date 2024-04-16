package monitoring

import (
	"github.com/common-fate/sdk/factoryconfig"
	"github.com/common-fate/sdk/gen/commonfate/factory/monitoring/v1alpha1/monitoringv1alpha1connect"
)

func NewFromConfig(cfg *factoryconfig.Context) monitoringv1alpha1connect.MonitoringServiceClient {
	return monitoringv1alpha1connect.NewMonitoringServiceClient(cfg.HTTPClient, cfg.BaseURL)
}
