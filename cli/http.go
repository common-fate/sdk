package cli

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/common-fate/clio"
	"github.com/common-fate/clio/clierr"
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/tokenstore"
)

// ErrorHandlingClient checks the response status code
// and creates an error if the API returns greater than 300.
type ErrorHandlingClient struct {
	Client     config.Doer
	LoginHint  string
	TokenStore *tokenstore.Storage
}

func (rd *ErrorHandlingClient) Do(req *http.Request) (*http.Response, error) {
	res, err := rd.Client.Do(req)
	var ne *url.Error
	if errors.As(err, &ne) && ne.Err == tokenstore.ErrNotFound {
		return nil, clierr.New(fmt.Sprintf("%s.\nTo get started with Common Fate, please run: '%s'", err, rd.LoginHint))

	}
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 300 {
		// response is ok
		return res, nil
	}

	// if we get here, the API has returned an error
	// surface this as a Go error so we don't need to handle it everywhere in our CLI codebase.
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return res, fmt.Errorf("reading error response body: %w", err)
	}
	bodyString := string(body)

	e := clierr.New(fmt.Sprintf("Common Fate API returned an error (code %v): %s", res.StatusCode, string(body)))
	if res.StatusCode == http.StatusUnauthorized {
		e.Messages = append(e.Messages, clierr.Infof("To log in to Common Fate, run: '%s'", rd.LoginHint))
	}

	if res.StatusCode == http.StatusBadRequest && strings.Contains(bodyString, "invalid_grant") {
		err = rd.TokenStore.Clear()
		if err != nil {
			return res, fmt.Errorf("error clearing cached Common Fate token: %w", err)
		}
		clio.Debugf("Cleared Common Fate cached token due to oauth2: cannot fetch token: 400, invalid_grant error")
		e.Messages = append(e.Messages, clierr.Infof("It looks like the above error was caused by an invalid authentication token. We have cleared the token from your keychain. To re-run the command, you'll need to authenticate again by running: '%s'", rd.LoginHint))
	}

	return res, e
}
