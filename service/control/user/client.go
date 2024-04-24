package attest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/user/v1alpha1/userv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) userv1alpha1connect.UserServiceClient {
	return userv1alpha1connect.NewUserServiceClient(cfg.HTTPClient, cfg.APIURL)
}
