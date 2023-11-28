package audit

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/access/v1alpha1/accessv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accessv1alpha1connect.AuditLogServiceClient {
	return accessv1alpha1connect.NewAuditLogServiceClient(cfg.HTTPClient, cfg.AccessURL)
}
