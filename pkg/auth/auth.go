package auth

import (
	"dotdev/nest/pkg/crud"
	"dotdev/nest/pkg/user"
)

const (
	DBTableOauth = "users_oauth"
)

// OAuth godoc
type OAuth struct {
	crud.Model
	crud.Timestampable

	UserPk uint64     `json:"-"`
	User   *user.User `json:"user" gorm:"references:pk;constraint:OnDelete:CASCADE;not null"`

	UniqueID string
	Provider string
}

func (OAuth) TableName() string {
	return "oauth"
}
