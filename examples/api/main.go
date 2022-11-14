package main

import (
	"os"

	"dotdev/nest/examples/api/config"
	"dotdev/nest/pkg/auth"
	"dotdev/nest/pkg/mailer"
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/nest/kernel"
	"dotdev/nest/pkg/nest/kernel/injector"
	"dotdev/nest/pkg/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	nest.LoadEnv()

	e := nest.New(
		kernel.New(),
		injector.Orm(),
		mailer.New(config.Hermes()),
		user.New(),
		auth.New(),
		// authcmd.NewRouter(),
	)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost", os.Getenv("CORS_ORIGIN")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.Serve())
}