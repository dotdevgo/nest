package user

import (
	"fmt"

	"github.com/dotdevgo/nest/pkg/crud"
	"github.com/go-playground/validator/v10"
	"github.com/goava/di"
)

type UserValidator struct {
	di.Inject
	Crud *UserCrud
}

// UniqueUsername implements validator.CustomTypeFunc
func (v UserValidator) UniqueUsername(fl validator.FieldLevel) bool {
	return v.uniqueField("username", fl)
}

// UniqueEmail implements validator.CustomTypeFunc
func (v UserValidator) UniqueEmail(fl validator.FieldLevel) bool {
	return v.uniqueField("email", fl)
}

// uniqueField godoc
func (v UserValidator) uniqueField(key string, fl validator.FieldLevel) bool {
	data, ok := fl.Parent().Interface().(crud.IModel)
	if !ok {
		return false
	}

	var counter int64 = 0
	sql := fmt.Sprintf("(%s.%s = ? AND %s.pk != ?)", DBTableUsers, key, DBTableUsers)
	if err := v.Crud.Stmt().
		Model(&User{}).
		Where(sql, fl.Field().String(), data.GetPk()).
		Count(&counter).
		Error; err != nil {
		return false
	}

	return counter == 0
}
