package app

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/defval/di"
// 	"gorm.io/gorm/clause"

// 	"dotdev/nest/crud"
// 	"dotdev/nest/nest"
// 	"dotdev/nest/orm"
// 	"dotdev/nest/paginator"
// 	"dotdev/nest/std/slices"
// )

// type AccountList []Account
// type AccountPaginator *paginator.Result[*AccountList]

// // AccountDTO godoc
// type AccountDTO struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name" form:"name" validate:"omitempty,ascii,min=3"`
// }

// // Account godoc
// type Account struct {
// 	crud.Entity

// 	Name     string            `json:"name"`
// 	Balances []*AccountBalance `json:"balances"`
// }

// // GetBalance godoc
// func (a *Account) GetBalance(currency string) *AccountBalance {
// 	if len(a.Balances) > 0 {
// 		balance := slices.Find[*AccountBalance](a.Balances, func(item *AccountBalance) bool {
// 			return item.Currency == currency
// 		})

// 		return balance
// 	}

// 	balance := &AccountBalance{
// 		AccountId: a.ID,
// 		Amount:    0,
// 		Currency:  currency,
// 	}

// 	a.Balances = append(a.Balances, balance)

// 	return balance
// }

// type AccountBalance struct {
// 	crud.Entity

// 	Amount   float64 `json:"amount" gorm:"type:decimal(60,30);"`
// 	Currency string  `json:"currency"`

// 	AccountId orm.BinaryUUID `json:"-" gorm:"type:binary(16);not null;"`
// 	Account   *Account       `json:"-" gorm:"references:id;constraint:OnDelete:CASCADE;not null"`
// }

// // accountRepository godoc
// type accountRepository struct {
// 	*crud.Repository[*Account]
// }

// // accountCrud godoc
// type accountCrud struct {
// 	di.Inject

// 	Repo *accountRepository
// }

// // NewAccount godoc
// func NewAccount() Account {
// 	account := Account{}
// 	account.SetId(crud.NewUUID().String())
// 	account.Balances = []*AccountBalance{}

// 	return account
// }

// // Find godoc
// func (c *accountCrud) Find(id string) (*Account, error) {
// 	var account = &Account{}
// 	account.SetId(id)

// 	if err := c.Repo.GetById(account); err != nil {
// 		return nil, err
// 	}

// 	return account, nil
// }

// // Paginate godoc
// func (c *accountCrud) Paginate(ctx nest.Context, options ...crud.Option) (AccountPaginator, error) {
// 	var result AccountList

// 	options = append(
// 		options,
// 		crud.WithPreload(clause.Associations),
// 		crud.WithCriteria(ctx.Request()),
// 	)

// 	stmt := c.Repo.CreateQuery(options...)

// 	return paginator.Paginate[*AccountList](stmt, &result, paginator.WithContext(ctx)...)
// }

// // Create godoc
// func (c *accountCrud) Create(accountDTO AccountDTO) (*Account, error) {
// 	account := NewAccount()
// 	account.Name = accountDTO.Name

// 	if len(accountDTO.Name) == 0 {
// 		account.Name = c.generateName(account.ID.String())
// 	}

// 	balance := account.GetBalance("USD")
// 	balance.Amount = 500.958458438585374537458423842834823

// 	if err := c.Repo.Create(&account).Error; err != nil {
// 		return nil, err
// 	}

// 	return &account, nil
// }

// // Update godoc
// func (c *accountCrud) Update(accountDTO AccountDTO) (*Account, error) {
// 	account, err := c.Find(accountDTO.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	account.Name = accountDTO.Name
// 	if err := c.Repo.Save(&account).Error; err != nil {
// 		return nil, err
// 	}

// 	return account, nil
// }

// func (c *accountCrud) generateName(id string) string {
// 	id = strings.ReplaceAll(id, "-", "")
// 	id = strings.ToUpper(id)

// 	return fmt.Sprintf("TRX%s", id)
// }

// /* Account Controller */
// type accountController struct {
// 	nest.Controller

// 	Repo *accountRepository
// 	Crud *accountCrud
// }

// // New godoc
// func (c accountController) New(w *nest.Kernel) {
// 	w.GET("/accounts", c.List)
// 	w.GET("/accounts/:id", c.Find)
// 	w.POST("/accounts/:id", c.Update)
// 	w.POST("/accounts", c.Create)
// }

// // List godoc
// func (c accountController) List(ctx nest.Context) error {
// 	p, err := c.Crud.Paginate(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.JSON(http.StatusOK, p)
// }

// // Find godoc
// func (c accountController) Find(ctx nest.Context) error {
// 	account, err := c.Crud.Find(ctx.Param("id"))
// 	if err != nil {
// 		return ctx.NoContent(http.StatusBadRequest)
// 	}

// 	return ctx.JSON(http.StatusOK, account)
// }

// // Create godoc
// func (c accountController) Create(ctx nest.Context) error {
// 	var input AccountDTO

// 	if err := ctx.Validate(&input); err != nil {
// 		return err
// 	}

// 	account, err := c.Crud.Create(input)
// 	if err != nil {
// 		return nest.NewHTTPError(http.StatusBadRequest, err)
// 	}

// 	return ctx.JSON(http.StatusOK, &account)
// }

// // Update godoc
// func (c accountController) Update(ctx nest.Context) error {
// 	var input = AccountDTO{
// 		ID: ctx.Param("id"),
// 	}

// 	if err := ctx.Validate(&input); err != nil {
// 		return err
// 	}

// 	account, err := c.Crud.Update(input)
// 	if err != nil {
// 		return nest.NewHTTPError(http.StatusBadRequest, err)
// 	}

// 	return ctx.JSON(http.StatusOK, &account)
// }

// /* Account Extension */
// type accountExtension struct {
// 	nest.Extension
// }

// // Boot godoc
// func (accountExtension) Boot(w *nest.Kernel) {
// 	w.Logger.Info("[AccountExtension] ==> BOOT")
// }

// func New() di.Option {
// 	return di.Options(
// 		orm.Mirgrate(&Account{}),
// 		orm.Mirgrate(&AccountBalance{}),
// 		di.Provide(func() *accountRepository {
// 			return &accountRepository{}
// 		}),
// 		di.Provide(func() *accountCrud {
// 			return &accountCrud{}
// 		}),
// 		nest.NewExtension(func() *accountExtension {
// 			return &accountExtension{}
// 		}),
// 		nest.NewController(func(repo *accountRepository) *accountController {
// 			return &accountController{}
// 		}),
// 	)
// }

// // *crud.Crud[*Account]
// // var accounts []Account
// // if err := c.Repo.FindAll(&accounts); err != nil {
// // 	return err
// // }
