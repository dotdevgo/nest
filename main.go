package main

import (
	"dotdev/nest/nest"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	w := nest.New()

	w.Use(middleware.Logger())
	w.Use(middleware.Recover())

	w.Logger.Fatal(w.Serve(":1323"))
}
