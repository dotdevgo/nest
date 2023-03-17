package user

import (
	"dotdev/nest/pkg/crud"
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/goava/di"
	"gorm.io/gorm"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&User{})
		}),
		di.Provide(crud.NewService[*User]),
		di.Provide(newUserCrud),
		di.Provide(func() *UserValidator {
			return &UserValidator{}
		}),
		di.Provide(func() *UserFactory {
			return &UserFactory{}
		}),
		nest.NewExtension(func() *userExt {
			return &userExt{}
		}),
	)
}

type userExt struct {
	nest.Extension
}

// Boot godoc
func (p userExt) Boot(w *nest.Kernel) error {
	w.InvokeFn(p.RegisterValidations)

	return nil
}

// RegisterValidations godoc
func (p userExt) RegisterValidations(w *nest.Kernel, uv *UserValidator, v *validator.Validate) {
	utils.NoErrorOrFatal(v.RegisterValidation("uniqueEmail", uv.UniqueEmail))
	utils.NoErrorOrFatal(v.RegisterValidation("uniqueUsername", uv.UniqueUsername))
}
