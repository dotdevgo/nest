package user

import (
	"dotdev/nest/pkg/crud"
	"fmt"
	"log"

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
	// sql := fmt.Sprintf("(%s.%s = ? AND %s.pk != ?)", DBTableUsers, key, DBTableUsers)
	if err := v.Crud.Stmt().
		Model(&User{}).
		Where(sql, args...). //fl.Field().String(), data.GetPk()).
		Count(&counter).
		Error; err != nil {
		return false
	}

	return counter == 0
}
