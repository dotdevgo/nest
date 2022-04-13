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
		RawAttributes *echo.Map      `gorm:"-" json:"attributes"`
		Attributes    datatypes.JSON `json:"-"`
	}

	Record interface {
		IsRecord()
	}

	IModel interface {
		GetID() uint
		GetUUID() string
	}
)

type Model struct {
	IModel `gorm:"-" json:"-"`
	ID     uint   `gorm:"primarykey" json:"-" meta:"getter;"`
	UUID   string `gorm:"type:varchar(255);uniqueIndex" json:"id" gqlgen:"id" meta:"getter;"`
}

func (Model) IsRecord() {}

// BeforeCreate godoc
func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == "" {
		u.UUID = uuid.New().String()
	}

	return
}

// BeforeSave godoc
func (u *Attributes) BeforeSave(tx *gorm.DB) (err error) {
	if err := u.initAttributes(); err != nil {
		return err
	}

	encoded, err := json.Marshal(u.RawAttributes)
	if err != nil {
		return err
	}

	u.Attributes = datatypes.JSON(encoded)

	return
}

// SetAttribute godoc
func (u *Attributes) SetAttribute(name string, value any) error {
	if err := u.initAttributes(); err != nil {
		return err
	}

	attr := *u.RawAttributes
	attr[name] = value
	u.RawAttributes = &attr

	return nil
}

// SetAttribute godoc
func (u *Attributes) DeleteAttribute(name string) error {
	if err := u.initAttributes(); err != nil {
		return err
	}

	attr := *u.RawAttributes
	delete(attr, name)
	u.RawAttributes = &attr

	return nil
}

// GetAttribute godoc
func (u *Attributes) GetAttribute(name string) any {
	if err := u.initAttributes(); err != nil {
		return nil
	}

	attr := *u.RawAttributes
	return attr[name]
}

// GetAttribute godoc
func (u *Attributes) initAttributes() (err error) {
	if nil != u.RawAttributes {
		return
	}

	var data *echo.Map = &echo.Map{}
	u.RawAttributes = data

	val, err := u.Attributes.Value()
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

// if nil == u.Attributes {
// 	u.Attributes = datatypes.JSON([]byte(`{"test":"row"}`))
// }
// val, err := u.Attributes.Value()
// var data echo.Map
// if err := json.Unmarshal([]byte(val.(string)), &data); err != nil {
// 	return err
// }
// data[name] = value
// encoded, err := json.Marshal(data)
// if err != nil {
// 	return err
// }
// u.Attributes = datatypes.JSON(encoded)

// fmt.Printf("%v %v %v", data, val.(string), err)
