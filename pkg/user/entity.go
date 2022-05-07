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
	DisplayName *string `json:"displayName"`
	Locale      *string `json:"locale"`
	Photo       string  `json:"photo" gorm:"null"`
	Bio         string  `json:"bio"`

	// TODO: remove
	// CountPublications  uint32 `json:"countPublications"`
	// CountFollowers     uint32 `json:"countFollowers"`
	// CountSubscriptions uint32 `json:"countSubscriptions"`

	IsVerified bool `json:"isVerified" gorm:"null"`
	IsDisabled bool `json:"isDisabled" gorm:"null"`

	Password []byte `json:"-" gorm:"not null;varchar(128);"`
}

func (u User) MarshalJSON() ([]byte, error) {
	displayName := u.DisplayName
	if displayName == nil {
		displayName = &u.Username
	}

	attributes, err := u.GetAttributes()
	if err != nil {
		return []byte{}, err
	}

	return json.Marshal(&struct {
		ID          string  `json:"id"`
		Email       string  `json:"email"`
		Username    string  `json:"username"`
		DisplayName *string `json:"displayName"`
		Bio         string  `json:"bio"`
		Photo       string  `json:"photo"`
		// CountPublications  uint32  `json:"countPublications"`
		// CountFollowers     uint32  `json:"countFollowers"`
		// CountSubscriptions uint32  `json:"countSubscriptions"`
		IsVerified bool `json:"isVerified"`
		IsDisabled bool `json:"isDisabled"`
		Attributes any  `json:"attributes"`
	}{
		ID:          u.ID,
		Email:       u.Email,
		Username:    u.Username,
		DisplayName: displayName,
		Bio:         u.Bio,
		Photo:       u.Photo,
		// CountPublications:  u.CountPublications,
		// CountFollowers:     u.CountFollowers,
		// CountSubscriptions: u.CountSubscriptions,
		IsVerified: u.IsVerified,
		IsDisabled: u.IsDisabled,
		Attributes: attributes,
	})
}
