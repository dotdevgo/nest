package auth

type (
	// AuthConfig stores the auth configuration
	AuthConfig struct {
		JwtSecret   string `env:"JWT_SECRET,default=secret"`
		SteamApiKey string `env:"STEAM_API_KEY,default=STEAM_API_KEY"`
	}
)
