package ingest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/granted/ingest/aws/v1alpha1/awsv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) awsv1alpha1connect.IngestServiceClient {
	return awsv1alpha1connect.NewIngestServiceClient(cfg.HTTPClient, cfg.APIURL)
}
