//go:generate metatag

package crud

import (
	"time"

	"gorm.io/gorm"
)

type (
	Model interface {
		GetID() string
	}

	Record interface {
		IsRecord()
	}

	SoftDeleteable struct {
		DeletedAt gorm.DeletedAt `json:"deletedAt" form:"deletedAt" gorm:"index"`
	}

	Timestampable struct {
		CreatedAt time.Time `json:"createdAt" gorm:"<-:create;"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	//CrudRepository struct{}
)

type Entity struct {
	Model `gorm:"-" json:"-"`
	ID    string `gorm:"type:varchar(255);primaryKey;uniqueIndex" json:"id" form:"id" gqlgen:"id"`
}

// GetID returns the value of UUID.
func (m *Entity) GetID() string {
	return m.ID
}

// GetPk() uint64
//Pk    uint64 `gorm:"primarykey" json:"-"`
// BeforeCreate godoc
// func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
// 	if m.ID == "" {
// 		m.ID = uuid.New().String()
// 	}

// 	return
// }

//func (Entity) IsRecord() {}

// GetPk returns the value of ID.
//func (m *Entity) GetPk() uint64 {
//	return m.Pk
//}
