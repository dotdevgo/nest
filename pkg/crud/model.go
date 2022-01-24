package crud

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	UUID      string         `gorm:"type:uuid;uniqueIndex" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// BeforeCreate godoc
func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}
