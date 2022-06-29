package user

import (
	"dotdev/nest/pkg/crud"
	"dotdev/nest/pkg/paginator"
)

type UserCrud struct {
	*crud.Crud[*User]
}

type UserList []*User
type UserPaginator *paginator.Result[*UserList]

// Paginate godoc
func (s UserCrud) Paginate(result interface{}, pagination []paginator.Option, options ...crud.Option) (UserPaginator, error) {
	var stmt = s.Stmt(options...)

	return paginator.Paginate[*UserList](stmt, result, pagination...)
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
