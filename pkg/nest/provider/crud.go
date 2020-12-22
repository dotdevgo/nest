package provider

import (
	"dotdev.io/pkg/crud"
	"github.com/goava/di"
)

func Crud() di.Option {
	return di.Options(
		di.Provide(crud.NewService),
	)
}