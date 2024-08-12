package ingest

import (
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/granted/ingest/v1alpha1/ingestv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) ingestv1alpha1connect.IngestServiceClient {
	return ingestv1alpha1connect.NewIngestServiceClient(cfg.HTTPClient, cfg.APIURL)
}
