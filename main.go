package main

import (
	"dotdev/events"
	"dotdev/extension"
	"dotdev/nest"
	"dotdev/orm"
	"dotdev/swagger"

	_ "dotdev/docs"

	"github.com/labstack/echo/v4/middleware"
)

// @title DotDev Golang packages
// @version 1.0
// @description Superset of golang useful packages
// @termsOfService https://dotdev.ltd/terms-of-service

// @contact.name DotDev API Support
// @contact.url http://dotdev.ltdo/support
// @contact.email me@dotdev.ltd

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	w := nest.New(
		extension.HealthCheck(),
		extension.Validator(),
		events.New(),
		orm.New(),
		swagger.New(),
	)

	w.Use(middleware.Logger())
	w.Use(middleware.Recover())

	w.Logger.Fatal(w.Serve(":1323"))
}
