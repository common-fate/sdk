package tokenstore

import (
	"errors"
	"sync"

	"golang.org/x/oauth2"
)

type InMemoryTokenStore struct {
	token *oauth2.Token
	mu    sync.Mutex
}

func NewInMemoryTokenStore() *InMemoryTokenStore {
	return &InMemoryTokenStore{}
}

func (ts *InMemoryTokenStore) Clear() error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.token = nil
	return nil
}

func (ts *InMemoryTokenStore) Save(token *oauth2.Token) error {
	if token == nil {
		return errors.New("token is nil")
	}

	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.token = token
	return nil
}

func (ts *InMemoryTokenStore) Token() (*oauth2.Token, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if ts.token == nil {
		return nil, errors.New("no token present")
	}

	// Return a copy of the token to prevent external modification
	tokenCopy := *ts.token
	return &tokenCopy, nil
}
