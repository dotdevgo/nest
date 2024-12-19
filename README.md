# Nest - Go Framework
```on top of labstack/echo@v4```


# main.go
```go
package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"dotdev/nest/kernel"
	"dotdev/nest/orm"
	"dotdev/nest/swagger"
)

// @title API Docs
// @version 1.0
// @description This is a sample Api server.
// @termsOfService https://dotdev.ltd/terms-of-service/

// @contact.name API Support
// @contact.url https://dotdev.ltd/support/
// @contact.email me@dotdev.ltd

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
func main() {
	e := nest.New(
		extension.HealthCheck(),
		extension.Validator(),
		events.New(),
		orm.New(),
		swagger.New(),
	)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost", os.Getenv("CORS_ORIGIN")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.Serve(":1323"))
}
```
