package log

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/control/log/v1alpha1/logv1alpha1connect"
)

type Client struct {
	cfg *config.Context
}

func New(cfg *config.Context) *Client {
	return &Client{cfg: cfg}
}

func (c Client) Evaluation() logv1alpha1connect.EvaluationServiceClient {
	return logv1alpha1connect.NewEvaluationServiceClient(c.cfg.HTTPClient, c.cfg.APIURL)
}
