package crud

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type (
	SoftDeleteable struct {
		DeletedAt gorm.DeletedAt `json:"deletedAt" form:"deletedAt" gorm:"index"`
	}

	Timestampable struct {
		CreatedAt time.Time `json:"createdAt" gorm:"<-:create;"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	Attributes struct {
		Attributes datatypes.JSON `json:"attributes" form:"attributes"`
	}
)

type Record interface {
	IsRecord()
}

type IModel interface {
	GetID() uint
	GetUUID() string
}

type Model struct {
	IModel    `gorm:"-" json:"-"`
	ID        uint           `gorm:"primarykey" json:"-"`
	UUID      string         `gorm:"type:uuid;uniqueIndex" json:"id"`
}

func (Model) IsRecord() {}

// GetID godoc
func (u *Model) GetID() uint {
	return u.ID
}

// GetUUID godoc
func (u *Model) GetUUID() string {
	return u.UUID
}

// BeforeCreate godoc
func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}
