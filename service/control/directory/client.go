package directory

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1/directoryv1alpha1connect"
)

type Client struct {
	cfg *config.Context
}

func New(cfg *config.Context) *Client {
	return &Client{cfg: cfg}
}

func (c Client) Evaluation() directoryv1alpha1connect.DirectoryServiceClient {
	return directoryv1alpha1connect.NewDirectoryServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}
