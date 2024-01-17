package provideraudittrail

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/audit/v1alpha1/auditv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) auditv1alpha1connect.ProviderAuditTrailServiceClient {
	return auditv1alpha1connect.NewProviderAuditTrailServiceClient(cfg.HTTPClient, cfg.APIURL)
}
