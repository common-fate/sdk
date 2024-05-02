package policyset

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

func NewFromConfig(cfg *config.Context) Client {
	url := cfg.APIURL

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
	rawClient := authzv1alpha1connect.NewPolicyServiceClient(&httpclient, url, connectOpts...)

	return Client{raw: rawClient}
}

var insecureTransport = &http2.Transport{
	AllowHTTP: true,
	DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
		return net.Dial(network, addr)
	},
}
