package schema

import (
	"crypto/tls"
	"net"
	"strings"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2"
)

type Opts struct {
	HTTPClient    connect.HTTPClient
	BaseURL       string
	ClientOptions []connect.ClientOption
}

func NewClient(opts Opts) authzv1alpha1connect.SchemaServiceClient {
	// client uses GRPC by default as the authz service does not support Buf Connect.
	connectOpts := []connect.ClientOption{connect.WithGRPC()}
	connectOpts = append(connectOpts, opts.ClientOptions...)

	return authzv1alpha1connect.NewSchemaServiceClient(opts.HTTPClient, opts.BaseURL, connectOpts...)
}

func NewFromConfig(cfg *config.Context) authzv1alpha1connect.SchemaServiceClient {

	connectOpts := []connect.ClientOption{connect.WithGRPC()}

	// dereference the client to avoid mutating it for all services
	// that share the config (this can cause HTTP2 issues in local dev)
	httpclient := cfg.HTTPClient
	if strings.HasPrefix(cfg.APIURL, "http://") {
		httpclient = cfg.HTTPClientBuilder(&oauth2.Transport{
			Source: cfg.TokenSource,
			Base:   insecureTransport,
		})
	}
	return authzv1alpha1connect.NewSchemaServiceClient(httpclient, cfg.APIURL, connectOpts...)
}

var insecureTransport = &http2.Transport{
	AllowHTTP: true,
	DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
		return net.Dial(network, addr)
	},
}
