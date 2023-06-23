package auth

import (
	"context"
	"dotdev/nest/pkg/nest"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware godoc
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()

			authCtx := NewContext(ctx.(nest.Context))

			user := authCtx.GetUser()

			if user == nil {
				return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			userCtx := context.WithValue(req.Context(), UserCtxKey, user)
			ctx.SetRequest(req.WithContext(userCtx))

			return next(ctx)
		}
	}
}
