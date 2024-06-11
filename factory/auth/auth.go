package auth

import (
	"net/http"
)

// Transport adds additional headers used by the Factory OIDC server
// to the outgoing HTTP requests.
type Transport struct {
	LicenceKey string
}

func (d *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if d.LicenceKey != "" {
		r.Header.Add("X-Common-Fate-Licence-Key", d.LicenceKey)
	}

	return http.DefaultTransport.RoundTrip(r)
}
