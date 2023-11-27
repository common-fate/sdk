package config

// IsConfigured returns true if a Common Fate config file
// exists and contains a non-empty current context.
//
// config.Load() is called to load the config file.
func IsConfigured() bool {
	cfg, err := load()
	if err != nil {
		return false
	}
	return cfg.CurrentContext != ""
}
