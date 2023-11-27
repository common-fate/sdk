package tokenstore

import (
	"github.com/99designs/keyring"
	"github.com/common-fate/clio"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Storage struct {
	keyring cfKeyring
	name    string
}

type Opts struct {
	Name string

	// Keyring is a custom keyring to use rather than
	// the default one, which is configured with
	// environment variables.
	Keyring keyring.Keyring
}

// WithKeyring specifies a custom keyring to use.
func WithKeyring(k keyring.Keyring) func(o *Opts) {
	return func(o *Opts) {
		o.Keyring = k
	}
}

// New creates a new token storage driver.
// The context is the authentication context to use.
// This is usually 'default' and in future can be
// expanded to allow CLI users to switch between
// separate Common Fate tenancies.
func New(opts Opts) Storage {
	return Storage{
		keyring: cfKeyring{keyring: opts.Keyring},
		name:    opts.Name,
	}
}

var (
	ErrNotFound = errors.New("auth token not found")
)

// Token returns the token.
func (s *Storage) Token() (*oauth2.Token, error) {
	var t oauth2.Token
	err := s.keyring.Retrieve(s.name, &t)
	if err != nil {
		clio.Debugf("error fetching token: %s", err)
		return &oauth2.Token{}, nil
	}

	return &t, nil
}

// Save the token
func (s *Storage) Save(token *oauth2.Token) error {
	return s.keyring.Store(s.name, token)
}

// Clear the token
func (s *Storage) Clear() error {
	return s.keyring.Clear(s.name)
}
