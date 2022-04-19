package user

import (
	"errors"

	"github.com/dotdevgo/nest/pkg/crud"
)

type UserCrud struct {
	*crud.Service[*User]
}

// FindByIdentity godoc
func (c UserCrud) FindByIdentity(identity string) (User, error) {
	var u User

	result := c.Stmt(ScopeByIdentity(identity)).First(&u)
	if err := result.Error; err != nil {
		return u, err
	}

	if 0 == u.Pk {
		return u, errors.New("Invalid identity")
	}

	return u, nil
}
