package kernel

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JwtClaims godoc
type JwtClaims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

// JwtMiddleware godoc
func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte("secret"),
	})
}
