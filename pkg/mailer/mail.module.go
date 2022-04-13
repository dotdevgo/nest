package mailer

import (
	"github.com/matcornic/hermes/v2"

	"github.com/goava/di"
)

// NewHermes godoc
func New(h *hermes.Hermes) di.Option {
	return di.Options(
		di.Provide(func() *hermes.Hermes {
			return h
		}),
	)
}
