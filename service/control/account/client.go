package attest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/account/accountv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) accountv1alpha1connect.AccountServiceClient {
	return accountv1alpha1connect.NewAccountServiceClient(cfg.HTTPClient, cfg.APIURL)
}
