package extension

import (
	"github.com/defval/di"
	"github.com/go-playground/validator/v10"
)

// Validator godoc
func Validator() di.Option {
	return di.Options(
		di.Provide(func() *validator.Validate {
			return validator.New()
		}),
	)
}
