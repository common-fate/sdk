package authz

import (
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
)

type Client struct {
	// Raw returns the underlying GRPC client.
	// It can be used to call methods that we don't have wrappers for yet.
	raw authzv1alpha1connect.AuthzServiceClient
}

type Opts struct {
	HTTPClient    connect.HTTPClient
	BaseURL       string
	ClientOptions []connect.ClientOption
	// AttrCacheExpiration defaults to 24h if not provided
	AttrCacheExpiration time.Duration
	// AttrCleanupInterval defaults to 48h if not provided
	AttrCleanupInterval time.Duration
}

func NewClient(opts Opts) Client {
	// client uses GRPC by default as the authz service does not support Buf Connect.
	connectOpts := []connect.ClientOption{connect.WithGRPC()}
	connectOpts = append(connectOpts, opts.ClientOptions...)

	rawClient := authzv1alpha1connect.NewAuthzServiceClient(opts.HTTPClient, opts.BaseURL, connectOpts...)

	return Client{raw: rawClient}
}
