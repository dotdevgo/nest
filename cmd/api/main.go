package main

import (
	"os"

	"github.com/dotdevgo/nest/cmd/api/config"
	authcmd "github.com/dotdevgo/nest/cmd/auth"
	"github.com/dotdevgo/nest/pkg/auth"
	"github.com/dotdevgo/nest/pkg/mailer"
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/dotdevgo/nest/pkg/nest/kernel"
	"github.com/dotdevgo/nest/pkg/nest/kernel/injector"
	"github.com/dotdevgo/nest/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	nest.LoadEnv()

	e := nest.New(
		kernel.New(),
		injector.NewOrm(),
		mailer.New(config.Hermes()),
		user.New(),
		auth.New(),
		authcmd.NewRouter(),
	)

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost", os.Getenv("CORS_ORIGIN")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.Serve())
}
