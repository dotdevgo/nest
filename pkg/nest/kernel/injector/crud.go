package injector

import (
	"github.com/dotdevgo/gosymfony/pkg/crud"
	"github.com/goava/di"
)

// Crud godoc
func Crud() di.Option {
	return di.Options(
		di.Provide(crud.NewService),
	)
}
