package auth

import (
	"context"
	"dotdev/nest/pkg/nest"

	"github.com/labstack/echo/v4"
)

// Middleware godoc
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := NewContext(ctx.(nest.Context))
			u := cc.User()

			if u != nil {
				authCtx := context.WithValue(ctx.Request().Context(), UserCtxKey, u)
				ctx.SetRequest(ctx.Request().WithContext(authCtx))
			}

			return next(ctx)
		}
	}
}