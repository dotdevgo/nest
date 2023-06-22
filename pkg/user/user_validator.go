package user

import (
	"dotdev/nest/pkg/crud"
	"fmt"
	"log"

	"github.com/defval/di"
	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	di.Inject
	Crud *UserCrud
}

// UniqueUsername implements validator.CustomTypeFunc
func (v UserValidator) UniqueUsername(fl validator.FieldLevel) bool {
	return v.validateUnique("username", fl)
}

// UniqueEmail implements validator.CustomTypeFunc
func (v UserValidator) UniqueEmail(fl validator.FieldLevel) bool {
	return v.validateUnique("email", fl)
}

// validateUnique godoc
// TODO: refactor
func (v UserValidator) validateUnique(key string, fl validator.FieldLevel) bool {
	data, ok := fl.Parent().Interface().(crud.Model)
	// if !ok {
	// 	// log.Fatalf("%v", fl.Parent().Interface())
	// 	return false
	// }

	var counter int64 = 0
	var sql string
	var args []any

	args = append(args, fl.Field().String())

	if ok && data.GetPk() > 0 {
		args = append(args, data.GetPk())
		sql = fmt.Sprintf("(%s.%s = ? AND %s.pk != ?)", DBTableUsers, key, DBTableUsers)
	} else {
		sql = fmt.Sprintf("(%s.%s = ?)", DBTableUsers, key)
	}

	log.Printf("ERROR %s %v", sql, args)

	if err := v.Crud.Stmt().
		Model(&User{}).
		Where(sql, args...). //fl.Field().String(), data.GetPk()).
		Count(&counter).
		Error; err != nil {
		return false
	}

	return counter == 0
}
