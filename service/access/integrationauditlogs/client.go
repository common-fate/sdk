package integrationauditlogs

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/access/v1alpha1/accessv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accessv1alpha1connect.IntegrationAuditLogServiceClient {
	return accessv1alpha1connect.NewIntegrationAuditLogServiceClient(cfg.HTTPClient, cfg.AccessURL)
}
