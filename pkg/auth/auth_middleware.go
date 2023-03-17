package auth

import (
	"context"
	"dotdev/nest/pkg/nest"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware godoc
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := Context(ctx.(nest.Context))
			u := c.User()
			if u != nil {
				authCtx := context.WithValue(ctx.Request().Context(), UserCtxKey, u)
				ctx.SetRequest(ctx.Request().WithContext(authCtx))
			}

			return next(ctx)
		}
	}
}
