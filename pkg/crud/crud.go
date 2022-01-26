package crud

import (
	"dotdev.io/pkg/nest"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

// NewService godoc
func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
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
	model := data.(IModel)
	if model.GetID() > 0 || model.GetUUID() != "" {
		return s.DB.Save(data).Error
	}

	return s.DB.Create(data).Error
}

// GetMany godoc
func (s *Service) GetMany(result interface{}, options ...Option) error {
	var stmt = s.newStmt(options...)

	return stmt.Find(result).Error
}

// newStmt godoc
func (s *Service) newStmt(options ...Option) *gorm.DB {
	var stmt = s.DB.Session(&gorm.Session{})

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}
