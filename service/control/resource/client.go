package attest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/resource/v1alpha1/resourcev1alpha1connect"
)

func NewFromConfig(cfg *config.Context) resourcev1alpha1connect.ResourceServiceClient {
	return resourcev1alpha1connect.NewResourceServiceClient(cfg.HTTPClient, cfg.APIURL)
}
