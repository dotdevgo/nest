package auth

type (
	// AuthConfig stores the auth configuration
	AuthConfig struct {
		JwtSecret string `env:"JWT_SECRET,default=secret"`
		// Steam
		SteamApiKey string `env:"STEAM_API_KEY,default=STEAM_API_KEY"`
		// Discord
		DiscordAppId  string `env:"DISCORD_APP_ID"`
		DiscordSecret string `env:"DISCORD_SECRET"`
	}
)
