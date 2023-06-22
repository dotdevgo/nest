package auth

import (
	"dotdev/nest/pkg/nest"

	"github.com/defval/di"
	"github.com/markbates/goth"
	"github.com/mustafaturan/bus/v3"
	"gorm.io/gorm"
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
		di.Provide(func() *AuthListener {
			return &AuthListener{}
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
	w.InvokeFn(p.RegisterMiddleware)
	w.InvokeFn(p.RegisterTopics)
	w.Invoke(p.RegisterProviders)

	return nil
}

// RegisterMiddleware godoc
func (p authExt) RegisterMiddleware(w *nest.Kernel, authConfig AuthConfig) {
	// TODO: refactor
	api := w.Api()
	api.Use(AuthMiddleware())
	api.Use(JwtMiddleware(authConfig))
}

// RegisterTopics godoc
func (p authExt) RegisterTopics(w *nest.Kernel, b *bus.Bus, h *AuthListener) {
	b.RegisterTopics(EventAuthRestore)
	b.RegisterHandler(EventAuthRestore, bus.Handler{
		Matcher: EventAuthRestore,
		Handle:  h.Restore,
	})

	b.RegisterTopics(EventAuthResetToken)
	b.RegisterHandler(EventAuthResetToken, bus.Handler{
		Matcher: EventAuthResetToken,
		Handle:  h.ResetToken,
	})

	b.RegisterTopics(EventAuthSignUp)
	b.RegisterHandler(EventAuthSignUp, bus.Handler{
		Matcher: EventAuthSignUp,
		Handle:  h.SignUp,
	})

	b.RegisterTopics(EventAuthConfirm)

	b.RegisterTopics(EventAuthResetEmail)
	b.RegisterHandler(EventAuthResetEmail, bus.Handler{
		Matcher: EventAuthResetEmail,
		Handle:  h.ResetEmail,
	})
}

// RegisterProviders godoc
func (p authExt) RegisterProviders(w *nest.Kernel, authConfig AuthConfig, providers []goth.Provider) {
	if nil == providers || len(providers) == 0 {
		return
	}

	goth.UseProviders(providers...)
}
