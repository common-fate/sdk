package loginflow

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/keyring"
	"github.com/common-fate/ciem/config"
	"github.com/common-fate/ciem/tokenstore"
	"github.com/common-fate/clio"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
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

	// Keyring optionally overrides the keyring that auth tokens are saved to.
	Keyring keyring.Keyring
}

type response struct {
	// Err is set if there was an error which
	// prevented the flow from completing
	Err   error
	Token (*oauth2.Token)
}

func (lf LoginFlow) Login(ctx context.Context) error {
	// oldCfg, err := cliconfig.Load()
	// if err != nil {
	// 	return err
	// }
	// oldDefaultContext := oldCfg.CurrentOrEmpty()

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
