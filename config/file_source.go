package config

import (
	"github.com/common-fate/clio"
	"go.uber.org/zap"
)

type FileSource struct {
	configFromFile Context
	loaded         bool
}

// Load config variables.
// The function must not set config variables if they are already set.
func (s *FileSource) Load(key Key) string {
	if !s.loaded {
		s.loaded = true
		configFromFile, err := load()
		if err != nil {
			clio.Debugw("error loading config from file", zap.Error(err))
			return ""
		}

		current, err := configFromFile.Current()
		if err != nil {
			clio.Debugw("error loading current config context from file", zap.Error(err))
			return ""
		}
		s.configFromFile = *current
	}

	switch key {
	case APIURLKey:
		return s.configFromFile.APIURL
	case AuthzURLKey:
		return s.configFromFile.AuthzURL
	case AccessURLKey:
		return s.configFromFile.AccessURL
	case OIDCIssuerKey:
		return s.configFromFile.OIDCIssuer
	case OIDCClientIDKey:
		return s.configFromFile.OIDCClientID
	case OIDCClientSecretKey:
		return s.configFromFile.OIDCClientSecret
	}

	return ""
}
