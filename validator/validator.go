package validator

import (
	"github.com/defval/di"
	"github.com/go-playground/validator/v10"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Provide(func() *validator.Validate {
			return validator.New()
		}),
	)
}
