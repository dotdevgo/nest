package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	dataform "github.com/dotdevgo/gosymfony/examples/app/dataform/handler/graphql"
	"github.com/dotdevgo/gosymfony/graphql/graph/generated"
	"github.com/goava/di"
)

type Resolver struct{
	generated.ResolverRoot
	di.Inject

	FormTemplateResolver *dataform.FormTemplateResolver
}
