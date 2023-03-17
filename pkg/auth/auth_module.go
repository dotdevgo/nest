package auth

import (
	"fmt"

	"dotdev/nest/pkg/logger"

	"github.com/joeshaw/envdecode"
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
		di.Provide(func() AuthConfig {
			var cfg AuthConfig
			if err := envdecode.StrictDecode(&cfg); err != nil {
				logger.Error(err)
			}
			// utils.NoErrorOrFatal(err)
			return cfg
		}),
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
		di.Provide(func() *module {
			return &module{}
		}, di.As(new(nest.Extension))),
	)
}

type module struct {
	nest.Extension
}

// Boot godoc
func (p module) Boot(w *nest.Kernel) error {
	p.RegisterTopics(w)
	p.RegisterAuthProviders(w)

	var authConfig AuthConfig
	w.ResolveFn(&authConfig)

	api := w.Secure()
	api.Use(JwtMiddleware(authConfig))
	api.Use(Middleware())

	return nil
}

// RegisterTopics godoc
func (p module) RegisterTopics(w *nest.Kernel) {
	var b *bus.Bus
	w.ResolveFn(&b)

	var h *AuthHooks
	w.ResolveFn(&h)

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
func (p module) RegisterAuthProviders(w *nest.Kernel) {
	var authConfig AuthConfig
	w.ResolveFn(&authConfig)

	callbackUrl := fmt.Sprintf("%s/auth/callback", w.Config.HTTP.Hostname)

	// TODO: refactor
	var arr []goth.Provider
	if authConfig.SteamApiKey != "" {
		steamProvider := steam.New(authConfig.SteamApiKey, fmt.Sprintf("%s/steam", callbackUrl))
		arr = append(arr, steamProvider)
		w.Logger.Info("[Auth] Register provider: \"steam\"")
	}

	if authConfig.DiscordAppId != "" {
		discordProvider := discord.New(authConfig.DiscordAppId, authConfig.DiscordSecret, fmt.Sprintf("%s/discord", callbackUrl))
		arr = append(arr, discordProvider)
		w.Logger.Info("[Auth] Register provider: \"discord\"")
	}

	goth.UseProviders(arr...)
}
