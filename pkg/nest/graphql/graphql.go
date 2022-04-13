package graphql

import (
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
)

// Provider godoc
func Provider() di.Option {
	return di.Options(
	// Graphql
	// di.Provide(func() *graph.Resolver {
	// 	return &graph.Resolver{}
	// }),
	)
}

// Router godoc
func Router(e *nest.EchoWrapper) {
	// GraphQL
	//var resolver graph.Resolver
	//e.ResolveFn(&resolver)
	//
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	//e.Any("/graphql", nest.WrapHandler(playground.Handler("GraphQL playground", "/graphql/query")))
	//e.Any("/graphql/query", nest.WrapHandler(srv))

}
