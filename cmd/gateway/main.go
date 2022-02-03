package main

import (
	"github.com/dotdevgo/gosymfony/examples/app/dataform"
	"github.com/dotdevgo/gosymfony/pkg/goutils"
	"github.com/dotdevgo/gosymfony/pkg/nest"
	"github.com/dotdevgo/gosymfony/pkg/nest/kernel"
	"github.com/dotdevgo/gosymfony/pkg/nest/kernel/injector"
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
		//di.Provide(func() *graph.Resolver {
		//	return &graph.Resolver{}
		//}),
		dataform.Provider(),
	)

	// GraphQL
	//var resolver graph.Resolver
	//e.ResolveFn(&resolver)
	//
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	//e.Any("/graphql", nest.WrapHandler(playground.Handler("GraphQL playground", "/graphql/query")))
	//e.Any("/graphql/query", nest.WrapHandler(srv))

	// Router
	//e.InvokeFn(dataform.Router)

	e.Logger.Fatal(e.Start(":1323"))
}
