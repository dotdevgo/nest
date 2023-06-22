package auth

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/steam"
	"gorm.io/gorm"

	"dotdev/nest/pkg/nest"

	"github.com/defval/di"
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
func (p authExt) OnStart(w *nest.Kernel) error {
	w.InvokeFn(p.registerMiddleware)
	w.InvokeFn(p.registerTopics)
	w.InvokeFn(p.registerProviders)

	return nil
}

// registerMiddleware godoc
func (p authExt) registerMiddleware(w *nest.Kernel, authConfig AuthConfig) {
	api := w.Secure()
	api.Use(JwtMiddleware(authConfig))
	api.Use(AuthMiddleware())
}

// registerTopics godoc
func (p authExt) registerTopics(w *nest.Kernel, b *bus.Bus, h *AuthHooks) {
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

// registerProviders godoc
func (p authExt) registerProviders(w *nest.Kernel, authConfig AuthConfig) {
	var providers []goth.Provider

	callbackUri := fmt.Sprintf("%s/auth/callback", w.Config.HTTP.Hostname)

	if authConfig.SteamApiKey != "" {
		steamProvider := steam.New(authConfig.SteamApiKey, fmt.Sprintf("%s/steam", callbackUri))
		providers = append(providers, steamProvider)
		w.Logger.Info("AUTH: Provider loaded ==> steam", authConfig.SteamApiKey)
	}

	if authConfig.DiscordSecret != "" {
		discordProvider := discord.New(authConfig.DiscordAppId, authConfig.DiscordSecret, fmt.Sprintf("%s/discord", callbackUri))
		providers = append(providers, discordProvider)
		w.Logger.Info("AUTH: Provider loaded ==> discord")
	}

	goth.UseProviders(providers...)
}
