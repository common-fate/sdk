package policyset

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"
	"golang.org/x/net/http2"
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
	url := cfg.AuthzURL
	if url == "" {
		url = cfg.APIURL
	}

	connectOpts := []connect.ClientOption{connect.WithGRPC()}

	var httpclient connect.HTTPClient
	if strings.HasPrefix(url, "http://") {
		httpclient = newInsecureClient()
	} else {
		httpclient = http.DefaultClient
	}
	rawClient := authzv1alpha1connect.NewPolicyServiceClient(httpclient, url, connectOpts...)

	return Client{raw: rawClient}
}

func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				// If you're also using this client for non-h2c traffic, you may want
				// to delegate to tls.Dial if the network isn't TCP or the addr isn't
				// in an allowlist.
				return net.Dial(network, addr)
			},
			// Don't forget timeouts!
		},
	}
}
