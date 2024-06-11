package loginflow

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/common-fate/clio"
	"github.com/common-fate/sdk/config"
	"github.com/common-fate/sdk/tokenstore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
)

func NewFromConfig(cfg *config.Context) LoginFlow {
	return LoginFlow{
		cfg: cfg,
	}
}

type LoginFlow struct {
	cfg *config.Context
}

type response struct {
	// Err is set if there was an error which
	// prevented the flow from completing
	Err   error
	Token (*oauth2.Token)
}
type myResponseWriter struct {
	header         http.Header
	statusCode     int
	body           []byte
	headerWritten  bool
	headerFlushed  bool
	originalWriter http.ResponseWriter
}

func (w *myResponseWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *myResponseWriter) Write(data []byte) (int, error) {
	if !w.headerWritten {
		w.WriteHeader(http.StatusOK)
	}
	w.body = append(w.body, data...)
	return len(data), nil
}

func (w *myResponseWriter) WriteHeader(statusCode int) {
	if w.headerFlushed {
		return
	}
	w.statusCode = statusCode
	w.headerWritten = true
}

func (lf LoginFlow) Login(ctx context.Context) error {

	authResponse := make(chan response)

	var g errgroup.Group

	// create random state variable for OIDC flow
	state := func() string {
		return uuid.New().String()
	}

	tokenWriter := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty) {
		authResponse <- response{
			Token: tokens.Token,
		}
		_, _ = w.Write([]byte("success! You can now close this window"))
	}

	r := chi.NewRouter()

	nonceErrorRetryAttempts := 0

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/auth/callback") {
				resWriter := &myResponseWriter{
					originalWriter: nil,
				}
				h.ServeHTTP(resWriter, r)

				res := string(resWriter.body)

				/*
					Sometimes the nonce will have a value when the auth library is not expecting it, typcally trying to login again fixes the issue.
					This is a related issue with Entra and the oidc client specifically see https://github.com/zitadel/oidc/issues/509
										So we have programmed this in and it will make one attempt to retry login before failing.
										To reproduce the none error, you go to the auth URL either custom like auth.myteam.com or the cognito url cf-auth-words.cognito.com and clear all the cookies
										Then try to login again, and you should get the nonce error
				*/
				if strings.Contains(res, "failed to exchange token: nonce does not match") {
					if nonceErrorRetryAttempts != 0 {
						// Write the original response after 1 retry attempt
						w.WriteHeader(http.StatusForbidden)
						_, _ = w.Write(resWriter.body)
						return
					}
					clio.Debugw("Caught a nonce validation error when exchanging code for token and will retry by redirecting to the login page", zap.Error(errors.New(res)))
					http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
					return
				}

				clio.Debugw("callback response body", "body", res)

				_, _ = w.Write(resWriter.body)
				return
			}
			h.ServeHTTP(w, r)
		})
	})
	r.Handle("/auth/callback", rp.CodeExchangeHandler(tokenWriter, lf.cfg.OIDCProvider))
	r.Handle("/login", rp.AuthURLHandler(state, lf.cfg.OIDCProvider, rp.WithPromptURLParam("Welcome back!")))
	server := &http.Server{
		Addr:    ":18900",
		Handler: r,
	}

	// run the auth server on localhost
	g.Go(func() error {
		clio.Debugw("starting HTTP server", "address", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		clio.Debugw("auth server closed")
		return nil
	})

	// read the returned ID token from Cognito
	g.Go(func() error {
		res := <-authResponse

		err := server.Shutdown(ctx)
		if err != nil {
			return err
		}

		// check that the auth flow didn't error out
		if res.Err != nil {
			return err
		}

		err = lf.cfg.TokenStore.Save(res.Token)
		if err != nil {
			return err
		}

		// reset the OAuth2.0 token source
		// NOTE(chrnorm): this should be moved to a method on the Context struct to avoid exposing internals
		src := &tokenstore.NotifyRefreshTokenSource{
			New:       lf.cfg.OIDCProvider.OAuthConfig().TokenSource(ctx, res.Token),
			T:         res.Token,
			SaveToken: lf.cfg.TokenStore.Save,
		}
		lf.cfg.TokenSource = src

		lf.cfg.HTTPClient = oauth2.NewClient(ctx, src)

		clio.Successf("Successfully logged in")

		return nil
	})

	// open the browser and read the token
	g.Go(func() error {
		url := "http://localhost:18900/login"
		clio.Infof("Opening your web browser to: %s", url)
		err := browser.OpenURL(url)
		if err != nil {
			clio.Errorf("error opening browser: %s", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (lf *LoginFlow) RefreshToken(ctx context.Context) (*oauth2.Token, error) {
	tok, err := lf.cfg.TokenStore.Token()
	if err != nil {
		return nil, err
	}

	c := lf.cfg.OIDCProvider.OAuthConfig()
	tok.Expiry = time.Now().Add(-time.Second * 10)
	src := tokenstore.NotifyRefreshTokenSource{
		New:       c.TokenSource(ctx, tok),
		T:         tok,
		SaveToken: lf.cfg.TokenStore.Save,
	}
	return src.Token()
}
