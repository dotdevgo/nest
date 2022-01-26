package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"dotdev.io/graphql/graph/generated"
	dataform "dotdev.io/internal/app/dataform/handler/graphql"
	"github.com/goava/di"
)

type Resolver struct{
	generated.ResolverRoot
	di.Inject

	FormTemplateResolver *dataform.FormTemplateResolver
}
