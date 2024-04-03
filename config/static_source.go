package config

type StaticSource struct {
	Result string
}

// Load config variables.
// The function must not set config variables if they are already set.
func (s StaticSource) Load(key Key) string {
	return s.Result
}
