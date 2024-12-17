package main

import (
	"dotdev/events"
	"dotdev/extension"
	"dotdev/nest"
	"dotdev/orm"
	"dotdev/swagger"
	tpl "dotdev/template"
	"html/template"
	"net/http"
	"os"

	_ "dotdev/docs"

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
	templates := template.Must(template.ParseGlob("*.html"))

	w := nest.New(
		extension.HealthCheck(),
		extension.Validator(),
		events.New(),
		tpl.New(templates),
		orm.New(),
		swagger.New(),
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

func home(c nest.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"Title": "DotDev",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
