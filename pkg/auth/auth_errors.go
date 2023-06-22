package auth

import "errors"

var (
	ErrorUnauthorized    = errors.New("unauthorized")
	ErrorInvalidPassword = errors.New("invalid body")
	ErrorUserDisabled    = errors.New("user is disabled")
	ErrorInvalidToken    = errors.New("invalid token")
)
