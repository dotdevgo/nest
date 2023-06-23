package config

import (
	"dotdev/nest/pkg/logger"

	"github.com/joeshaw/envdecode"
)

type AppAuthConfig struct {
	// Discord
	// DiscordAppId  string `env:"DISCORD_APP_ID"`
	// DiscordSecret string `env:"DISCORD_SECRET"`
	// Steam
	// SteamApiKey string `env:"STEAM_API_KEY"`
}

// NewAuthConfig godoc
func NewAuthConfig() AppAuthConfig {
	var cfg AppAuthConfig

	if err := envdecode.StrictDecode(&cfg); err != nil {
		logger.Error(err)
	}

	return cfg
}

// di.Provide(config.NewAuthConfig),
// di.Provide(func(authConfig config.AppAuthConfig) goth.Provider {
// 	return discord.New(authConfig.DiscordAppId, authConfig.DiscordSecret, "/discord")
// }),
// di.Provide(func(authConfig config.AppAuthConfig) goth.Provider {
// 	return steam.New(authConfig.SteamApiKey, "/steam")
// }),
