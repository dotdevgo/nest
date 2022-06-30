package auth

import "errors"

var (
	ErrorInvalidPassword = errors.New("Invalid body")
	ErrorUserDisabled    = errors.New("User is disabled")
)
