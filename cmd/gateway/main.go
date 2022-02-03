package main

import (
	"dotdev.io/examples/app/dataform"
	"dotdev.io/graphql/graph/generated"
	graph "dotdev.io/graphql/graph/resolver"
	"dotdev.io/pkg/goutils"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/kernel"
	"dotdev.io/pkg/nest/kernel/injector"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/goava/di"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"os"
)

func main() {
	goutils.NoErrorOrFatal(godotenv.Load(".env"))

	e := nest.New(
		injector.Orm(mysql.Open(os.Getenv("DATABASE")), nil),
		kernel.Provider(),
		// Graphql
		di.Provide(func() *graph.Resolver {
			return &graph.Resolver{}
		}),
		dataform.Provider(),
	)

	// GraphQL
	var resolver graph.Resolver
	e.ResolveFn(&resolver)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	e.Any("/graphql", nest.WrapHandler(playground.Handler("GraphQL playground", "/graphql/query")))
	e.Any("/graphql/query", nest.WrapHandler(srv))

	// Router
	e.InvokeFn(dataform.Router)

	e.Logger.Fatal(e.Start(":1323"))
}
