package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// URLSource configures Common Fate from a deployment URL.
// The URL should be a config.json which contains discovery variables
// for a particular deployment.
type URLSource struct {
	DeploymentURL *url.URL
	configJSON    *URLConfigJSON
}

// Load config variables.
func (u *URLSource) Load(key Key) (string, error) {
	if u.configJSON == nil {
		err := u.fetchConfig()
		if err != nil {
			return "", err
		}
	}

	switch key {
	case APIURLKey:
		return u.configJSON.APIURL, nil
	case OIDCClientIDKey:
		return u.configJSON.CliOAuthClientId, nil
	case OIDCIssuerKey:
		return u.configJSON.OAauthIssuer, nil
	case AccessURLKey:
		return u.configJSON.AccessAPIURL, nil
	}

	return "", nil
}

func (u *URLSource) fetchConfig() error {
	res, err := http.DefaultClient.Get(u.DeploymentURL.String())
	if err != nil {
		return fmt.Errorf("error fetching config from %s: %w", u.DeploymentURL.String(), err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var cfg URLConfigJSON
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return err
	}
	u.configJSON = &cfg
	return nil
}

type URLConfigJSON struct {
	CliOAuthClientId string `json:"cliOAuthClientId"`
	OAauthIssuer     string `json:"oauthIssuer"`
	APIURL           string `json:"apiUrl"`
	AccessAPIURL     string `json:"accessApiUrl"`
}
