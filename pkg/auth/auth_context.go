package auth

import (
	"dotdev/nest/pkg/user"

	"dotdev/nest/pkg/nest"

	"github.com/golang-jwt/jwt"

	"context"
)

// AuthContext godoc
type AuthContext struct {
	nest.Context
}

// User godoc
func (c AuthContext) User() *user.User {
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

// Context godoc
func Context(ctx nest.Context) *AuthContext {
	cc := &AuthContext{Context: ctx}
	return cc
}

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *user.User {
	raw, _ := ctx.Value(UserCtxKey).(*user.User)
	return raw
}
