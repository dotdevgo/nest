package auth

import (
	"dotdev/nest/pkg/logger"

	"github.com/joeshaw/envdecode"
)

type (
	// AuthConfig stores the auth configuration
	AuthConfig struct {
		JwtExpire uint   `env:"JWT_EXPIRE,default=36000"`
		JwtSecret string `env:"JWT_SECRET,default=secret"`
		// CallbackUrl string `env:"OAUTH_CALLBACK_URL"`
	}
)

// NewAuthConfig godoc
func NewAuthConfig() AuthConfig {
	var cfg AuthConfig

	if err := envdecode.StrictDecode(&cfg); err != nil {
		logger.Panic(err)
	}

	return cfg
}
