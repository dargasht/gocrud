package gocrud

type Config struct {
	authSecret string
	language   string
}

type ConfigOption func(*Config)

func WithAuthSecret(authSecret string) ConfigOption {
	return func(c *Config) {
		c.authSecret = authSecret
	}
}

func WithLanguage(language string) ConfigOption {
	return func(c *Config) {
		c.language = language
	}
}

// SetConfig sets the config
// Look at the main file in the example
func SetConfig(options ...ConfigOption) *Config {
	c := Config{
		authSecret: "secret",
		language:   "fa",
	}
	for _, option := range options {
		option(&c)
	}
	return &c
}

// GoCRUDConfig is the default config
// You should set this to your needs in your main file
// Look at the main file in the example
var GoCRUDConfig = SetConfig()
