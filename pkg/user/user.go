package user

import (
	"encoding/json"

	"dotdev/nest/pkg/crud"
)

const (
	DBTableUsers = "users"
)

type Email string

// User godoc
type User struct {
	crud.Model
	crud.Timestampable
	crud.Attributes

	Email       string  `json:"email" gorm:"not null;index:uniqueEmail,unique"`
	Username    string  `json:"username" gorm:"not null;index:uniqueUsername,unique"`
	DisplayName string  `json:"displayName"`
	Locale      *string `json:"locale"`

	IsVerified bool `json:"isVerified" gorm:"null"`
	IsDisabled bool `json:"isDisabled" gorm:"null"`

	Password []byte `json:"-" gorm:"not null;varchar(128);"`
}

func (u User) MarshalJSON() ([]byte, error) {
	// attributes, err := u.GetAttributes()
	// if err != nil {
	// 	return []byte{}, err
	// }

	return json.Marshal(&struct {
		ID          string `json:"id"`
		Email       string `json:"email"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		IsVerified  bool   `json:"isVerified"`
		IsDisabled  bool   `json:"isDisabled"`
		//Attributes  crud.JSON `json:"attributes"`
	}{
		ID:          u.ID,
		Email:       u.Email,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		IsVerified:  u.IsVerified,
		IsDisabled:  u.IsDisabled,
		//Attributes:  attributes,
	})
}
