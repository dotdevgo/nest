package main

import (
	"dotdev.io/graphql/graph/generated"
	graph "dotdev.io/graphql/graph/resolver"
	"dotdev.io/internal/app/dataform"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/provider"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goava/di"
	"gorm.io/driver/sqlite"
)

func main() {
	e := nest.New(
		provider.Orm(sqlite.Open("datastore.db"), nil),
		provider.Validator(),
		provider.Crud(),
		dataform.Provider(),
		di.Provide(func() *graph.Resolver {
			return &graph.Resolver{}
		}),
	)

	// GraphQL
	var resolver graph.Resolver
	e.ResolveFn(&resolver)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	e.Any("/graphql", nest.WrapHandler(playground.Handler("GraphQL playground", "/graphql/query")))
	e.Any("/graphql/query", nest.WrapHandler(srv))

	// Apps
	e.InvokeFn(dataform.Router)

	e.Logger.Fatal(e.Start(":1323"))
}
