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
		di.Provide(NewUserCrud),
		di.Provide(func() *UserValidator {
			return &UserValidator{}
		}),
		di.Provide(func() *UserFactory {
			return &UserFactory{}
		}),
		di.Provide(func() *UserModule {
			return &UserModule{}
		}, di.As(new(nest.ContainerModule))),
	)
}

// NewUserCrud godoc
func NewUserCrud(c *crud.Crud[*User]) *UserCrud {
	return &UserCrud{
		Crud: c,
	}
}

// UserModule godoc
type UserModule struct {
	nest.ContainerModule
}

// Boot godoc
func (p UserModule) Boot(w *nest.Kernel) error {
	p.RegisterValidations(w)

	return nil
}

// RegisterValidations godoc
func (p UserModule) RegisterValidations(w *nest.Kernel) {
	var uv *UserValidator
	w.ResolveFn(&uv)

	var v *validator.Validate
	w.ResolveFn(&v)

	utils.NoErrorOrFatal(v.RegisterValidation("uniqueEmail", uv.UniqueEmail))
	utils.NoErrorOrFatal(v.RegisterValidation("uniqueUsername", uv.UniqueUsername))
}
