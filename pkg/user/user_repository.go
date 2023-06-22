package user

import (
	"dotdev/nest/pkg/orm"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserRepository struct {
	orm.Repository[*User]
	*gorm.DB
}

func (r UserRepository) FinByConfirmToken(token string) *User {
	var u User

	result := r.CreateQueryBuilder().
		Find(&u, datatypes.JSONQuery("attributes").
			Equals(token, "confirmToken"))

	if result.Error != nil || u.Pk <= 0 {
		return nil
	}

	return &u
}
