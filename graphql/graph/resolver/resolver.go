package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	dataform "dotdev.io/examples/app/dataform/handler/graphql"
	"dotdev.io/graphql/graph/generated"
	"github.com/goava/di"
)

type Resolver struct{
	generated.ResolverRoot
	di.Inject

	FormTemplateResolver *dataform.FormTemplateResolver
}
