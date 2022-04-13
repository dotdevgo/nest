package main

import (
	"os"

	authcmd "github.com/dotdevgo/nest/cmd/auth"
	"github.com/dotdevgo/nest/pkg/auth"
	"github.com/dotdevgo/nest/pkg/mailer"
	nest "github.com/dotdevgo/nest/pkg/nest"
	"github.com/dotdevgo/nest/pkg/nest/kernel"
	"github.com/dotdevgo/nest/pkg/nest/kernel/injector"
	"github.com/dotdevgo/nest/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	hr "github.com/matcornic/hermes/v2"
)

func main() {
	nest.LoadEnv()

	e := nest.New(
		kernel.New(),
		injector.NewOrm(),
		mailer.New(hermes()),
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

	e.Logger.Fatal(e.Start())
}

func hermes() *hr.Hermes {
	return &hr.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hr.Product{
			Name:      "DotDevio",
			Link:      "https://dotdevio.com/",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2022 DotDevio. All rights reserved",
		},
	}
}
