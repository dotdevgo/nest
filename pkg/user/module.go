package user

import (
	"github.com/dotdevgo/nest/pkg/crud"
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/dotdevgo/nest/pkg/utils"
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
		di.Provide(func() *UserProvider {
			return &UserProvider{}
		}, di.As(new(nest.ServiceProvider))),
	)
}

// NewUserCrud godoc
func NewUserCrud(c *crud.Service[*User]) *UserCrud {
	return &UserCrud{
		Service: c,
	}
}

// UserProvider godoc
type UserProvider struct {
	nest.ServiceProvider
}

// Boot godoc
func (p UserProvider) Boot(w *nest.Kernel) error {
	p.RegisterValidations(w)

	return nil
}

// RegisterValidations godoc
func (p UserProvider) RegisterValidations(w *nest.Kernel) {
	var uv *UserValidator
	w.ResolveFn(&uv)

	var v *validator.Validate
	w.ResolveFn(&v)

	utils.NoErrorOrFatal(v.RegisterValidation("uniqueEmail", uv.UniqueEmail))
	utils.NoErrorOrFatal(v.RegisterValidation("uniqueUsername", uv.UniqueUsername))
}
