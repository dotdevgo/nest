//go:generate metatag

package orm

import (
	"dotdev/logger"
	"time"

	"github.com/gofrs/uuid/v5"

	"gorm.io/gorm"
)

type (
	Model interface {
		GetId() string
		// SetId(id string) error
	}

	Entity struct {
		Model `json:"-" gorm:"-"`
		ID    BinaryUUID `json:"id" form:"id" gqlgen:"id" gorm:"type:binary(16);primaryKey;notNull;"`
	}

	SoftDeleteable struct {
		DeletedAt gorm.DeletedAt `json:"deletedAt" form:"deletedAt" gorm:"index"`
	}

	Timestampable struct {
		CreatedAt time.Time `json:"createdAt" gorm:"<-:create;"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)

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

	m.ID = BinaryUUID(uuid)

	return nil
}

// BeforeCreate godoc
func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == BinaryUUID(uuid.Nil) {
		m.ID = UUIDToBinary(NewUUID().String())
	}

	return
}

// NewUUID godoc
func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()

	logger.PanicOnError(err)

	return id
}
