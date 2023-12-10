package policyset

import (
	"github.com/bufbuild/connect-go"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
)

type Client struct {
	raw authzv1alpha1connect.PolicyServiceClient
}

type Opts struct {
	HTTPClient    connect.HTTPClient
	BaseURL       string
	ClientOptions []connect.ClientOption
}

func NewClient(opts Opts) Client {
	// client uses GRPC by default as the authz service does not support Buf Connect.
	connectOpts := []connect.ClientOption{connect.WithGRPC()}
	connectOpts = append(connectOpts, opts.ClientOptions...)

	rawClient := authzv1alpha1connect.NewPolicyServiceClient(opts.HTTPClient, opts.BaseURL, connectOpts...)

	return Client{raw: rawClient}
}
