package main

import (
	"os"

	"dotdev/nest/pkg/kernel"
	"dotdev/nest/pkg/kernel/extension"
	"dotdev/nest/pkg/nest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "dotdev/nest/cmd/api/docs"

	swagger "github.com/swaggo/echo-swagger"
)

// @title API Docs
// @version 1.0
// @description This is a sample Api server.
// @termsOfService http://dotdev.io/terms/

// @contact.name API Support
// @contact.url http://www.dotdev.io/support
// @contact.email support@dotdev.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
func main() {
	e := nest.New(
		kernel.New(),
		extension.Orm(),
	)

	e.Echo.GET("/docs/*", swagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost", os.Getenv("CORS_ORIGIN")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.Serve(":1323"))
}
