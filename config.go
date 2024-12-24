package gocrud

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config for GoCRUD
type Config struct {
	appName    string
	authSecret string
	language   string
	otpApiKey  string
	s3Client   *s3.Client
	bucketName string
}

type ConfigOption func(*Config)

// SetConfig sets the config
// Look at the main file in the example
func SetConfig(options ...ConfigOption) *Config {
	c := Config{
		appName:    "gocrud",
		authSecret: "secret",
		language:   "fa",
		otpApiKey:  "your_otp_api_key",
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

// Options
func WithAppName(appName string) ConfigOption {
	return func(c *Config) {
		c.appName = appName
	}
}
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

func WithOtpApiKey(otpApiKey string) ConfigOption {
	return func(c *Config) {
		c.otpApiKey = otpApiKey
	}
}

func WithS3Client(s3Client *s3.Client, bucketName string) ConfigOption {
	return func(c *Config) {
		c.s3Client = s3Client
		c.bucketName = bucketName
	}
}
