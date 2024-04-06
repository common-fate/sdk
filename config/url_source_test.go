package config

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestURLSource_Load(t *testing.T) {

	tests := []struct {
		key          Key
		name         string
		mockResponse string
		want         string
		wantErr      bool
	}{
		{
			name: "ok",
			key:  OIDCClientIDKey,
			mockResponse: `{
				"oauthClientId": "XXX",
				"cliOAuthClientId": "YYY",
				"oauthAuthority": "https://cognito-idp.ap-southeast-2.amazonaws.com/ap-southeast-2_ABCD/",
				"oauthIssuer": "https://cognito-idp.ap-southeast-2.amazonaws.com/ap-southeast-2_ABCD",
				"apiUrl": "https://commonfate.example.com",
				"accessApiUrl": "https://commonfate.example.com",
				"authzGraphApiUrl": "https://commonfate.example.com/graph",
				"teamName": "Common Fate",
				"iconUrl": "",
				"releaseTag": "sha-48b2449"
			  }`,
			want: "YYY",
		},
		{
			// should just return an empty string for the client secret
			name: "ok_for_client_secret_which_is_not_served_by_config_json",
			key:  OIDCClientSecretKey,
			mockResponse: `{
				"oauthClientId": "XXX",
				"cliOAuthClientId": "YYY",
				"oauthAuthority": "https://cognito-idp.ap-southeast-2.amazonaws.com/ap-southeast-2_ABCD/",
				"oauthIssuer": "https://cognito-idp.ap-southeast-2.amazonaws.com/ap-southeast-2_ABCD",
				"apiUrl": "https://commonfate.example.com",
				"accessApiUrl": "https://commonfate.example.com",
				"authzGraphApiUrl": "https://commonfate.example.com/graph",
				"teamName": "Common Fate",
				"iconUrl": "",
				"releaseTag": "sha-48b2449"
			  }`,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			depURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}

			u := &URLSource{
				DeploymentURL: depURL,
			}
			got, err := u.Load(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLSource.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLSource.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
