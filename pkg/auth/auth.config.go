package auth

type (
	// AuthConfig stores the auth configuration
	AuthConfig struct {
		SteamApiKey string `env:"STEAM_API_KEY"`
	}
)
