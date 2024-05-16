package tokenstore

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/keyring"
	"github.com/common-fate/clio"
	"golang.org/x/oauth2"
)

type WindowsStorage struct {
	keyring Keyring
	name    string
}

type WindowsOpts struct {
	Name string

	// Keyring is a custom keyring to use rather than
	// the default one, which is configured with
	// environment variables.
	Keyring keyring.Keyring
}

// New creates a new token storage driver.
// The context is the authentication context to use.
// This is usually 'default' and in future can be
// expanded to allow CLI users to switch between
// separate Common Fate tenancies.
func NewWindows(opts Opts) WindowsStorage {
	return WindowsStorage{
		keyring: Keyring{keyring: opts.Keyring},
		name:    opts.Name,
	}
}

// Token returns the token.
func (s *WindowsStorage) Token() (*oauth2.Token, error) {
	var key []byte
	clio.Debugf("attempting to fetch token: %s", s.name)
	err := s.keyring.Retrieve(s.name, &key)
	if err != nil {
		clio.Debugf("error fetching token: %s", err)
		return &oauth2.Token{}, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		clio.Debugf("could not create new cipher: %v", err)
		return &oauth2.Token{}, nil
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		clio.Debugf("could not create new cipher: %v", err)
		return &oauth2.Token{}, nil
	}

	cd, err := os.UserConfigDir()
	if err != nil {
		clio.Debugf("could not find user dir", err)
		return &oauth2.Token{}, nil
	}
	dir := filepath.Join(cd, "commonfate")
	// Trim the URL components from the name but keep it readable
	suffix := strings.TrimPrefix(s.name, "http://")
	suffix = strings.TrimPrefix(suffix, "https://")
	suffix = strings.ReplaceAll(suffix, "/", "_")
	encryptedTokenFile := filepath.Join(dir, "auth_token_"+suffix)

	ciphertext, err := os.ReadFile(encryptedTokenFile)
	if err != nil {
		clio.Debugf("could not read encrypted Token File : %v", err)
		return &oauth2.Token{}, nil
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		clio.Debugf("ciphertext too short : %v", err)
		return &oauth2.Token{}, nil
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		clio.Debugf("could not decrypt file : %v", err)
		return &oauth2.Token{}, nil
	}

	var token oauth2.Token
	if err := json.Unmarshal(plaintext, &token); err != nil {
		clio.Debugf("could not unmarshal token: %v", err)
		return &oauth2.Token{}, nil
	}

	return &token, nil
}

// Save the token
func (s *WindowsStorage) Save(token *oauth2.Token) error {
	// create a new encryption key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create new cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// encrypt the token and store it in the user app data dir
	cd, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(cd, "commonfate")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating %s: %w", dir, err)
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("error marshalling token: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, tokenJSON, nil)

	// Trim the URL components from the name but keep it readable
	suffix := strings.TrimPrefix(s.name, "http://")
	suffix = strings.TrimPrefix(suffix, "https://")
	suffix = strings.ReplaceAll(suffix, "/", "_")
	encryptedTokenFile := filepath.Join(dir, "auth_token_"+suffix)

	err = os.WriteFile(encryptedTokenFile, ciphertext, 0600)
	if err != nil {
		return err
	}

	return s.keyring.Store(s.name, key)
}

// Clear the token
func (s *WindowsStorage) Clear() error {
	return s.keyring.Clear(s.name)
}
