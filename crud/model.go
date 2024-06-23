//go:generate metatag

package crud

import (
	"dotdev/orm"
	"time"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type (
	Model interface {
		GetId() string
		// SetId(id string) error
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
	Model `json:"-" gorm:"-"`
	ID    orm.BinaryUUID `json:"id" form:"id" gqlgen:"id" gorm:"type:binary(16);primaryKey;notNull;"`
}

// GetId returns the value of UUID.
func (m Entity) GetId() string {
	return m.ID.String()
}

// SetId godoc
func (m *Entity) SetId(id string) error {
	uuid, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	m.ID = orm.BinaryUUID(uuid)

	return nil
}

// BeforeCreate godoc
func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == orm.BinaryUUID(uuid.Nil) {
		m.ID = orm.UUIDToBinary(NewUUID().String())
	}

	return
}

// GetPk() uint64
//Pk    uint64 `gorm:"primarykey" json:"-"`

//func (Entity) IsRecord() {}

// GetPk returns the value of ID.
//func (m *Entity) GetPk() uint64 {
//	return m.Pk
//}
