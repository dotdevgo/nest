//go:generate metatag

package crud

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	SoftDeleteable struct {
		DeletedAt gorm.DeletedAt `json:"deletedAt" form:"deletedAt" gorm:"index"`
	}

	Timestampable struct {
		CreatedAt time.Time `json:"createdAt" gorm:"<-:create;"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	Record interface {
		IsRecord()
	}

	Model interface {
		// GetPk() uint64

		GetID() string
	}

	CrudRepository struct{}
)

type Entity struct {
	Model `gorm:"-" json:"-"`
	ID    string `gorm:"type:varchar(255);uniqueIndex" json:"id" form:"id" gqlgen:"id"`
	//Pk    uint64 `gorm:"primarykey" json:"-"`
}

// GetID returns the value of UUID.
func (m *Entity) GetID() string {
	return m.ID
}

// BeforeCreate godoc
func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return
}

//func (Entity) IsRecord() {}

// GetPk returns the value of ID.
//func (m *Entity) GetPk() uint64 {
//	return m.Pk
//}
