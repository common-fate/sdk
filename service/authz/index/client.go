package index

import (
	"github.com/bufbuild/connect-go"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
)

func NewClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) authzv1alpha1connect.IndexServiceClient {
	// client ALWAYS uses GRPC as Rust does not support Buf Connect.
	opts = append(opts, connect.WithGRPC())
	return authzv1alpha1connect.NewIndexServiceClient(httpClient, baseURL, opts...)
}
