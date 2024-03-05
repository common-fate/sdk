package diagnostic

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/deployment/diagnostic/v1alpha1/diagnosticv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) diagnosticv1alpha1connect.DiagnosticServiceClient {
	return diagnosticv1alpha1connect.NewDiagnosticServiceClient(cfg.HTTPClient, cfg.APIURL)
}
