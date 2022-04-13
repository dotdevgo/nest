package auth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewJwtClaims godoc
func NewJwtClaims(identity string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Subject:   identity,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
}

// JwtMiddleware godoc
func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &jwt.StandardClaims{},
		// TODO: token secret
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}

// JwtClaims godoc
// type JwtClaims struct {
// 	Identity string `json:"identity"`
// 	jwt.StandardClaims
// }

// return &JwtClaims{
// 	username,
// 	jwt.StandardClaims{
// 		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
// 	},
// }
