package authcmd

import (
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
)

// NewRouter godoc
func NewRouter() di.Option {
	return di.Options(
		di.Provide(func() *AuthController {
			return &AuthController{}
		}, di.As(new(nest.Controller))),
	)
}
