package injector

import (
	"github.com/dotdevgo/nest/pkg/crud"
	"github.com/goava/di"
)

// Crud godoc
func Crud() di.Option {
	return di.Options(
		di.Provide(crud.NewService),
	)
}
