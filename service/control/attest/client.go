package attest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1/attestv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) attestv1alpha1connect.AttestServiceClient {
	return attestv1alpha1connect.NewAttestServiceClient(cfg.HTTPClient, cfg.APIURL)
}
