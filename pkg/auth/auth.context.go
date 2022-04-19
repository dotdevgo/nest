package auth

import (
	"github.com/dotdevgo/nest/pkg/user"

	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/golang-jwt/jwt"
)

// AuthContext godoc
type AuthContext struct {
	nest.Context
}

// User godoc
func (c AuthContext) User() *user.User {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	var crud *user.UserCrud
	if err := c.Resolve(&crud); err != nil {
		return nil
	}

	var u user.User
	if err := crud.Find(&u, claims.Subject); err != nil {
		return nil
	}

	return &u
}

// NewContext godoc
func NewContext(ctx nest.Context) *AuthContext {
	cc := &AuthContext{Context: ctx}
	return cc
}
