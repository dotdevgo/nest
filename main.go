package main

import (
	"dotdev/extension"
	"dotdev/nest"
	"dotdev/orm"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	w := nest.New(
		extension.HealthCheck(),
		extension.Validator(),
		orm.New(),
	)

	w.Use(middleware.Logger())
	w.Use(middleware.Recover())

	w.Logger.Fatal(w.Serve(":1323"))
}
