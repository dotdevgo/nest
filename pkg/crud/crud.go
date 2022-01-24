package crud

import (
	paginator "dotdev.io/pkg/gorm-paginator"
	"dotdev.io/pkg/goutils"
	"dotdev.io/pkg/nest"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

// NewService godoc
func NewService(db *gorm.DB) Service {
	return Service{DB: db}
	// .Session(&gorm.Session{NewDB: true})
}

// IsValid godoc
func (s *Service) IsValid(c nest.Context, input interface{}) error {
	if err := c.Bind(input); err != nil {
		return err //echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return err //echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

// Save godoc
func (s *Service) Save(data interface{}) error {
	var model = new(Model)
	goutils.Copy(model, data)

	if model.ID > 0 || model.UUID != "" {
		return s.DB.Save(data).Error
	}

	return s.DB.Create(data).Error
}

// GetMany godoc
func (s *Service) GetMany(result interface{}, options ...Option) error {
	var stmt = s.newStmt(options...)

	return stmt.Find(result).Error
}

// Paginate godoc
func (s *Service) Paginate(result interface{}, pagination []paginator.Option, options ...Option) (*paginator.Result, error) {
	var stmt = s.newStmt(options...)

	return paginator.Paginate(stmt, result, pagination...)
}

// newStmt godoc
func (s *Service) newStmt(options ...Option) *gorm.DB {
	var stmt = s.DB

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}
