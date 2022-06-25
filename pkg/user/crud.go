package user

import (
	"dotdev/nest/pkg/crud"
)

type UserCrud struct {
	*crud.Crud[*User]
}

// FindByIdentity godoc
func (c UserCrud) FindByIdentity(identity string) (User, error) {
	var u User

	result := c.Stmt(ScopeByIdentity(identity)).First(&u)
	if err := result.Error; err != nil {
		return u, err
	}

	if 0 == u.Pk {
		return u, ErrorInvalidIdentity
	}

	return u, nil
}
