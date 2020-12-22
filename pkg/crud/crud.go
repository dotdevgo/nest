package crud

import (
	paginator "dotdev.io/pkg/gorm-paginator"
	"dotdev.io/pkg/nest"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	UUID      string         `gorm:"type:uuid;uniqueIndex" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) IsValid(c nest.Context, input interface{}) error {
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *Service) Create(data interface{}) error {
	return s.DB.Create(data).Error
}

func (s *Service) Paginate(c nest.Context, result interface{}) (*paginator.Result, error) {
	options := []paginator.Option{
		paginator.WithPage(2),
		paginator.WithLimit(10),
		//paginator.WithOrder("ID DESC"),
	}

	return paginator.Paginate(s.DB, result, options...)
}

//s.DB.Find(result)
//func (s *Service) throwError(c nest.Context, err error, statusCode int) error {
//	return c.JSON(statusCode, map[string]interface{}{"error":err.Error()})
//}

//func Validate(validate *validator.Validate, c nest.Context, input interface{}) error {
//	if err := validate.Struct(input); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error":err.Error()})
//	}
//
//	return nil
//}
