package main

import (
	"dotdev.io/internal/app/dataform"
	"dotdev.io/internal/app/dataform/entity"
	"dotdev.io/pkg/crud"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/provider"
	"github.com/goava/di"
	"gorm.io/driver/sqlite"
)

func main() {
	e := nest.New(
		provider.Orm(sqlite.Open("datastore.db"), &provider.OrmConfig{
			Entities: []interface{}{
				&entity.FormTemplate{},
			},
		}),
		provider.Validator(),
		di.Options(
			di.Provide(crud.NewService),
		),
	)

	dataform.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}
