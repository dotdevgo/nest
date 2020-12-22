package main

import (
	"dotdev.io/internal/app/dataform"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/provider"
	"gorm.io/driver/sqlite"
)

func main() {
	e := nest.New(
		provider.Orm(sqlite.Open("datastore.db"), dataform.OrmConfig),
		provider.Validator(),
		provider.Crud(),
	)

	dataform.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}
