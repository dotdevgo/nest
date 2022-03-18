package crud

import (
	nest "github.com/dotdevgo/nest/pkg/core"
	"github.com/labstack/gommon/log"
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
	log.Printf("ID: %v UUID: %v", model.GetID(), model.GetUUID())
	if model.GetID() > 0 || model.GetUUID() != "" {
		return s.DB.Save(data).Error
	}

	return s.DB.Create(data).Error
}

// Find godoc
func (s *Service) Find(result interface{}, id interface{}, options ...Option) error {
	stmt := s.newStmt(options...)

	return stmt.Scopes(ScopeById(result, id)).First(result).Error
}

// GetMany godoc
func (s *Service) GetMany(result interface{}, options ...Option) error {
	var stmt = s.newStmt(options...)

	return stmt.Find(result).Error
}

// newStmt godoc
func (s *Service) newStmt(options ...Option) *gorm.DB {
	var stmt = s.DB.Session(&gorm.Session{NewDB: true})

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}
