package auth

import (
	"fmt"
	"log"

	"github.com/dotdevgo/nest/pkg/goutils"
	"github.com/dotdevgo/nest/pkg/user"
	"github.com/joeshaw/envdecode"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"

	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Provide(func() AuthConfig {
			var cfg AuthConfig
			err := envdecode.StrictDecode(&cfg)
			goutils.NoErrorOrFatal(err)
			return cfg
		}),
		di.Provide(func() *AuthService {
			return &AuthService{}
		}),
		di.Provide(func() *AuthHooks {
			return &AuthHooks{}
		}),
		di.Provide(func() *AuthMailer {
			return &AuthMailer{}
		}),
		di.Provide(func() *AuthProvider {
			return &AuthProvider{}
		}, di.As(new(nest.ServiceProvider))),
	)
}

// AuthProvider godoc
type AuthProvider struct {
	nest.ServiceProvider
}

// Boot godoc
func (p AuthProvider) Boot(w *nest.Kernel) error {
	p.RegisterTopics(w)
	p.RegisterAuthProviders(w)

	api := w.ApiGroup()
	api.Use(JwtMiddleware())

	return nil
}

// RegisterTopics godoc
func (p AuthProvider) RegisterTopics(w *nest.Kernel) {
	var b *bus.Bus
	w.ResolveFn(&b)

	var h *AuthHooks
	w.ResolveFn(&h)

	b.RegisterTopics(user.EventUserRestore)
	b.RegisterHandler(user.EventUserRestore, h.EventRestore())

	b.RegisterTopics(user.EventUserResetToken)
	b.RegisterHandler(user.EventUserResetToken, h.EventResetToken())
}

// RegisterAuthProviders godoc
func (p AuthProvider) RegisterAuthProviders(w *nest.Kernel) {
	var authConfig AuthConfig
	w.ResolveFn(&authConfig)

	var arr []goth.Provider
	if authConfig.SteamApiKey != "" {
		arr = append(arr, steam.New(authConfig.SteamApiKey, fmt.Sprintf("%s/auth/callback/steam", w.Config.HTTP.Hostname)))
		log.Printf("[OAuth] Registered \"steam\" provider")
	}
	goth.UseProviders(arr...)
}
