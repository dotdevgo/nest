package auth

import "errors"

var (
	ErrorUnauthorized    = errors.New("Unauthorized")
	ErrorInvalidPassword = errors.New("Invalid body")
	ErrorUserDisabled    = errors.New("User is disabled")
)
