package auth

import (
	"dotdev/nest/pkg/user"

	"dotdev/nest/pkg/nest"

	"github.com/golang-jwt/jwt"
)

// фuthContext godoc
type authContext struct {
	nest.Context
}

// GetUser godoc
func (c authContext) GetUser() *user.User {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil
	}

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
func NewContext(ctx nest.Context) *authContext {
	cc := &authContext{Context: ctx}
	return cc
}

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// GetUser finds the user from the context. REQUIRES Middleware to have run.
func GetUser(ctx nest.Context) *user.User {
	c := ctx.Request().Context()

	raw, _ := c.Value(UserCtxKey).(*user.User)

	return raw
}
