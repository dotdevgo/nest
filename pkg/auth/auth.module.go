package auth

import (
	"github.com/dotdevgo/nest/pkg/user"

	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// NewModule godoc
func New() di.Option {
	return di.Options(
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
func (p *AuthProvider) Boot(w *nest.EchoWrapper) error {
	p.RegisterTopics(w)

	api := w.ApiGroup()
	api.Use(JwtMiddleware())

	return nil
}

// RegisterTopics godoc
func (p *AuthProvider) RegisterTopics(w *nest.EchoWrapper) {
	var b *bus.Bus
	w.ResolveFn(&b)

	var h *AuthHooks
	w.ResolveFn(&h)

	b.RegisterTopics(user.EventUserRestore)
	b.RegisterHandler(user.EventUserRestore, h.EventRestore())

	b.RegisterTopics(user.EventUserResetToken)
	b.RegisterHandler(user.EventUserResetToken, h.EventResetToken())
}
