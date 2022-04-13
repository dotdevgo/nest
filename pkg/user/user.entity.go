package user

import (
	"encoding/json"

	"github.com/dotdevgo/nest/pkg/crud"
)

const (
	DBTableUsers = "users"
)

const (
	AttributeResetToken   = "reset_token"
	AttributeConfirmToken = "confirmToken"
)

type Email string

// User godoc
type User struct {
	crud.Model
	crud.Timestampable
	crud.Attributes

	// todo: unique indexes
	Email       string  `json:"email" gorm:"not null;index:uniqueEmail,unique"`
	Username    string  `json:"username" gorm:"not null;index:uniqueUsername,unique"`
	DisplayName *string `json:"displayName"`
	Locale      *string `json:"locale"`
	Bio         string  `json:"bio"`
	Photo       string  `json:"photo" gorm:"null"`

	CountPublications  uint32 `json:"countPublications"`
	CountFollowers     uint32 `json:"countFollowers"`
	CountSubscriptions uint32 `json:"countSubscriptions"`

	IsVerified bool `json:"isVerified" gorm:"null"`
	IsDisabled bool `json:"isDisabled" gorm:"null"`

	Password []byte `json:"-" gorm:"not null;varchar(128);"`
}

func (u User) MarshalJSON() ([]byte, error) {
	displayName := u.DisplayName
	if displayName == nil {
		displayName = &u.Username
	}

	return json.Marshal(&struct {
		UUID               string  `json:"id"`
		Email              string  `json:"email"`
		Username           string  `json:"username"`
		DisplayName        *string `json:"displayName"`
		Bio                string  `json:"bio"`
		Photo              string  `json:"photo"`
		CountPublications  uint32  `json:"countPublications"`
		CountFollowers     uint32  `json:"countFollowers"`
		CountSubscriptions uint32  `json:"countSubscriptions"`
		IsVerified         bool    `json:"isVerified"`
		IsDisabled         bool    `json:"isDisabled"`
	}{
		UUID:               u.UUID,
		Email:              u.Email,
		Username:           u.Username,
		DisplayName:        displayName,
		Bio:                u.Bio,
		Photo:              u.Photo,
		CountPublications:  u.CountPublications,
		CountFollowers:     u.CountFollowers,
		CountSubscriptions: u.CountSubscriptions,
		IsVerified:         u.IsVerified,
		IsDisabled:         u.IsDisabled,
	})
}
