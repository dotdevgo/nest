package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewJwtClaims godoc
func NewJwtClaims(identity string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Subject: identity,
		// TODO: refactor expires at based on option/argument
		// ExpiresAt: time.Now().Add(15).Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
}

// JwtMiddleware godoc
func JwtMiddleware(config AuthConfig) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte(config.JwtSecret),
		// TODO: refactor
		Skipper: func(ctx echo.Context) bool {
			token := ctx.Request().Header.Get("Authorization")
			return token == ""
		},
	})
}
