package auth

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/steam"
	"gorm.io/gorm"

	"dotdev/nest/pkg/nest"

	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Provide(NewAuthConfig),
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&OAuth{})
		}),
		di.Provide(func() *AuthManager {
			return &AuthManager{}
		}),
		di.Provide(func() *AuthHooks {
			return &AuthHooks{}
		}),
		di.Provide(func() *AuthMailer {
			return &AuthMailer{}
		}),
		nest.NewExtension(func() *authExt {
			return &authExt{}
		}),
	)
}

type authExt struct {
	nest.Extension
}

// Boot godoc
func (p authExt) Boot(w *nest.Kernel) error {
	w.InvokeFn(p.RegisterTopics)
	w.InvokeFn(p.RegisterAuthProviders)

	var authConfig AuthConfig
	w.ResolveFn(&authConfig)

	api := w.Secure()
	api.Use(JwtMiddleware(authConfig))
	api.Use(Middleware())

	return nil
}

// RegisterTopics godoc
func (p authExt) RegisterTopics(w *nest.Kernel, b *bus.Bus, h *AuthHooks) {
	b.RegisterTopics(EventUserRestore)
	b.RegisterHandler(EventUserRestore, bus.Handler{
		Matcher: EventUserRestore,
		Handle:  h.Restore,
	})

	b.RegisterTopics(EventUserResetToken)
	b.RegisterHandler(EventUserResetToken, bus.Handler{
		Matcher: EventUserResetToken,
		Handle:  h.ResetToken,
	})

	b.RegisterTopics(EventUserSignUp)
	b.RegisterHandler(EventUserSignUp, bus.Handler{
		Matcher: EventUserSignUp,
		Handle:  h.SignUp,
	})

	b.RegisterTopics(EventUserConfirm)

	b.RegisterTopics(EventUserResetEmail)
	b.RegisterHandler(EventUserResetEmail, bus.Handler{
		Matcher: EventUserResetEmail,
		Handle:  h.ResetEmail,
	})
}

// RegisterAuthProviders godoc
func (p authExt) RegisterAuthProviders(w *nest.Kernel, authConfig AuthConfig) {
	// var authConfig AuthConfig
	// w.ResolveFn(&authConfig)

	callbackUri := fmt.Sprintf("%s/auth/callback", w.Config.HTTP.Hostname)

	var providers []goth.Provider
	if authConfig.SteamApiKey != "" {
		steamProvider := steam.New(authConfig.SteamApiKey, fmt.Sprintf("%s/steam", callbackUri))
		providers = append(providers, steamProvider)
		w.Logger.Info("[Auth] Register provider: \"steam\"")
	}

	if authConfig.DiscordAppId != "" {
		discordProvider := discord.New(authConfig.DiscordAppId, authConfig.DiscordSecret, fmt.Sprintf("%s/discord", callbackUri))
		providers = append(providers, discordProvider)
		w.Logger.Info("[Auth] Register provider: \"discord\"")
	}

	goth.UseProviders(providers...)
}
