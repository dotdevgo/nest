package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/goava/di"
)

// Validator godoc
func Validator() di.Option {
	return di.Options(
		di.Provide(func() *validator.Validate {
			return validator.New()
		}),
	)
}
