package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(pass string) (password []byte, err error) {
	password, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}

	return password, err
}
