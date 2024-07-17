package directory

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1/directoryv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) directoryv1alpha1connect.DirectoryServiceClient {
	return directoryv1alpha1connect.NewDirectoryServiceClient(cfg.HTTPClient, cfg.APIURL)
}
