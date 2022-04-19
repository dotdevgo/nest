//go:generate metatag

package crud

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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

	Attributes struct {
		RawAttributes echo.Map       `gorm:"-" json:"attributes"`
		Attributes    datatypes.JSON `json:"-"`
	}

	Record interface {
		IsRecord()
	}

	IModel interface {
		GetPk() uint64
		GetID() string
	}
)

type Model struct {
	IModel `gorm:"-" json:"-"`
	Pk     uint64 `gorm:"primarykey" json:"-" meta:"getter;"`
	ID     string `gorm:"type:varchar(255);uniqueIndex" json:"id" gqlgen:"id" meta:"getter;"`
}

func (Model) IsRecord() {}

// BeforeCreate godoc
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}

	return
}

// BeforeSave godoc
func (m *Attributes) BeforeSave(tx *gorm.DB) (err error) {
	if err := m.parseAttributes(); err != nil {
		return err
	}

	encoded, err := json.Marshal(m.RawAttributes)
	if err != nil {
		return err
	}

	m.Attributes = datatypes.JSON(encoded)

	return
}

// SetAttribute godoc
func (m *Attributes) SetAttribute(name string, value any) error {
	if err := m.parseAttributes(); err != nil {
		return err
	}

	attr := m.RawAttributes
	attr[name] = value
	m.RawAttributes = attr

	return nil
}

// GetAttribute godoc
func (m *Attributes) GetAttribute(name string) any {
	if err := m.parseAttributes(); err != nil {
		return nil
	}

	attr := m.RawAttributes
	return attr[name]
}

// SetAttribute godoc
func (m *Attributes) DeleteAttribute(name string) error {
	if err := m.parseAttributes(); err != nil {
		return err
	}

	attr := m.RawAttributes
	delete(attr, name)
	m.RawAttributes = attr

	return nil
}

// GetAttribute godoc
func (m *Attributes) parseAttributes() (err error) {
	if nil != m.RawAttributes {
		return
	}

	var data echo.Map = echo.Map{}
	m.RawAttributes = data

	val, err := m.Attributes.Value()
	if err != nil {
		return err
	}

	if nil == val {
		return
	}

	if err := json.Unmarshal([]byte(val.(string)), &data); err != nil {
		return err
	}

	return
}

// if nil == m.Attributes {
// 	m.Attributes = datatypes.JSON([]byte(`{"test":"row"}`))
// }
// val, err := m.Attributes.Value()
// var data echo.Map
// if err := json.Unmarshal([]byte(val.(string)), &data); err != nil {
// 	return err
// }
// data[name] = value
// encoded, err := json.Marshal(data)
// if err != nil {
// 	return err
// }
// m.Attributes = datatypes.JSON(encoded)

// fmt.Printf("%v %v %v", data, val.(string), err)
