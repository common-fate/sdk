package authz

import (
	"crypto/tls"
	"net"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2"
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

func NewFromConfig(cfg *config.Context) Client {
	url := cfg.AuthzURL
	if url == "" {
		url = cfg.APIURL
	}

	connectOpts := []connect.ClientOption{connect.WithGRPC()}

	// dereference the client to avoid mutating it for all services
	// that share the config (this can cause HTTP2 issues in local dev)
	httpclient := *cfg.HTTPClient
	if strings.HasPrefix(url, "http://") {
		httpclient.Transport = &oauth2.Transport{
			Source: cfg.TokenSource,
			Base:   insecureTransport,
		}
	}
	rawClient := authzv1alpha1connect.NewAuthzServiceClient(&httpclient, url, connectOpts...)

	return Client{raw: rawClient}
}

var insecureTransport = &http2.Transport{
	AllowHTTP: true,
	DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
		return net.Dial(network, addr)
	},
}
