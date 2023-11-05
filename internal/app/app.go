package app

import (
	"net/http"

	"github.com/defval/di"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"dotdev/nest/pkg/crud"
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/orm"
	"dotdev/nest/pkg/paginator"
)

type AccountList []*Account
type AccountPaginator *paginator.Result[*AccountList]

// Account godoc
type Account struct {
	crud.Entity

	Name string `json:"name"`
}

// NewAccount godoc
func NewAccount() Account {
	account := Account{}
	account.ID = uuid.New().String()

	return account
}

// AccountDTO godoc
type AccountDTO struct {
	Id   string
	Name string `json:"name" form:"name" validate:"omitempty,ascii,min=3"`
}

type AccountRepository struct {
	*crud.Repository[*Account]
}

type AccountCrud struct {
	di.Inject

	Repo *AccountRepository
}

// Paginate godoc
func (c *AccountRepository) Paginate(
	pagination []paginator.Option,
	options ...crud.Option,
) (AccountPaginator, error) {
	var result AccountList

	options = append(
		options,
		crud.WithPreload(clause.Associations),
	)

	stmt := c.CreateQuery(options...)

	return paginator.Paginate[*AccountList](stmt, &result, pagination...)
}

// Create godoc
func (c *AccountCrud) Create(accountDTO AccountDTO) (*Account, error) {
	account := NewAccount()
	account.Name = accountDTO.Name

	if err := c.Repo.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

/* Account Controller */
type accountController struct {
	nest.Controller

	Repo *AccountRepository
	Crud *AccountCrud
}

// New godoc
func (c accountController) New(w *nest.Kernel) {
	w.GET("/accounts", c.List)
	w.GET("/account/:id", c.Get)
	w.POST("/accounts", c.Create)
}

// List godoc
func (c accountController) List(ctx nest.Context) error {
	p, err := c.Repo.Paginate(paginator.WithContext(ctx))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, p)
}

// Get godoc
func (c accountController) Get(ctx nest.Context) error {
	account := Account{}

	stmt := c.Repo.First(&account, "id = ?", ctx.Param("id"))
	if err := stmt.Error; err != nil {
		return err
	}

	if len(account.ID) == 0 {
		return ctx.NotFound()
	}

	return ctx.JSON(http.StatusOK, &account)
}

// Create godoc
func (c accountController) Create(ctx nest.Context) error {
	var input AccountDTO

	if err := ctx.Validate(&input); err != nil {
		return err
	}

	account, err := c.Crud.Create(input)

	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, &account)
}

/* Account Extension */
type accountExtension struct {
	nest.Extension
}

// Boot godoc
func (accountExtension) Boot(w *nest.Kernel) {
	w.Logger.Info("[AccountExtension] ==> BOOT")
}

func New() di.Option {
	return di.Options(
		orm.Mirgrate[*Account](&Account{}),
		di.Provide(func() *AccountRepository {
			return &AccountRepository{}
		}),
		di.Provide(func() *AccountCrud {
			return &AccountCrud{}
		}),
		nest.NewExtension(func() *accountExtension {
			return &accountExtension{}
		}),
		nest.NewController(func(repo *AccountRepository) *accountController {
			return &accountController{}
		}),
	)
}

// *crud.Crud[*Account]
// var accounts []Account
// if err := c.Repo.FindAll(&accounts); err != nil {
// 	return err
// }
