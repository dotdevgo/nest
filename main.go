// +templ generate
package main

import (
	"dotdev/events"
	"dotdev/logger"
	"dotdev/nest"
	"dotdev/orm"
	"dotdev/swagger"
	"dotdev/template"
	"dotdev/validator"
	html "html/template"
	"net/http"
	"os"

	_ "dotdev/docs"

	"github.com/defval/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title DotDev Golang packages
// @version 1.0
// @description Superset of golang useful packages
// @termsOfService https://dotdev.ltd/terms-of-service/

// @contact.name DotDev API Support
// @contact.url https://dotdev.ltdo/support/
// @contact.email me@dotdev.ltd

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
func main() {
	templates := html.Must(html.ParseGlob("*.html"))

	container, err := di.New()
	if err != nil {
		logger.Fatal(err)
	}

	config := []nest.Option{
		nest.UseContainer(container),
	}

	w := nest.NewWithConfig(
		config,
		orm.New(),
		events.New(),
		swagger.New(),
		validator.New(),
		template.New(templates),
	)

	w.Use(middleware.Logger())
	w.Use(middleware.Recover())

	w.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost", os.Getenv("CORS_ORIGIN")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	w.GET("/", home)

	w.Logger.Fatal(w.Serve(":1323"))
}

// TODO:
// https://www.reddit.com/r/golang/comments/17d12wk/using_echo_with_ahtempl/
func home(c nest.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"Title": "DotDev",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
