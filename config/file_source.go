package config

import "errors"

// ErrConfigFileNotFound is returned if the config file (~/.cf/config by default)
// does not exist.
var ErrConfigFileNotFound = errors.New("cf config file does not exist (~/.cf/config)")

type FileSource struct {
	configFromFile Context
	loaded         bool
}

// Load config variables.
func (s *FileSource) Load(key Key) (string, error) {
	if !s.loaded {
		s.loaded = true
		configFromFile, err := load()
		if err != nil {
			return "", err
		}

		current, err := configFromFile.Current()
		if err != nil {
			return "", err
		}
		s.configFromFile = *current
	}

	switch key {
	case APIURLKey:
		return s.configFromFile.APIURL, nil
	case AuthzURLKey:
		return s.configFromFile.AuthzURL, nil
	case AccessURLKey:
		return s.configFromFile.AccessURL, nil
	case OIDCIssuerKey:
		return s.configFromFile.OIDCIssuer, nil
	case OIDCClientIDKey:
		return s.configFromFile.OIDCClientID, nil
	case OIDCClientSecretKey:
		return s.configFromFile.OIDCClientSecret, nil
	case NameKey:
		return s.configFromFile.name, nil
	}

	return "", nil
}
