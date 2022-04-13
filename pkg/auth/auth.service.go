package auth

import (
	"context"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"

	"github.com/dgrijalva/jwt-go"
	"github.com/dotdevgo/nest/pkg/goutils"
	"github.com/dotdevgo/nest/pkg/user"
	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// AuthService godoc
type AuthService struct {
	di.Inject
	Crud *user.UserCrud
	Bus  *bus.Bus
}

// Validate godoc
func (c *AuthService) Validate(input *SignInDto) (*user.User, error) {
	u, err := c.Crud.FindByIdentity(input.Identity)
	if err != nil {
		return nil, err
	}

	if u.IsDisabled {
		return nil, errors.New("User is disabled")
	}

	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(input.Password)); err != nil {
		return nil, ErrInvalidPassword
	}

	return u, nil
}

// SignUp godoc
func (c *AuthService) SignUp(input *SignUpDto) (*user.User, error) {
	var u = new(user.User)
	goutils.Copy(u, input)

	// Defaults
	u.CountPublications = 0
	u.CountFollowers = 0
	u.CountSubscriptions = 0

	u.IsVerified = false
	u.IsDisabled = false

	u.SetAttribute(user.AttributeConfirmToken, goutils.RandomToken())

	// Password
	pass, err := c.hashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	u.Password = pass

	if err := c.Crud.Save(u); err != nil {
		return nil, err
	}

	if err := c.Bus.Emit(context.Background(), user.EventUserSignUp, u); err != nil {
		return nil, err
	}

	return u, nil
}

// Confirm godoc
func (c *AuthService) Confirm(token string) error {
	var u user.User

	result := c.Crud.NewStmt().
		Find(&u, datatypes.JSONQuery("attributes").
			Equals(token, user.AttributeConfirmToken))

	if result.Error != nil || u.ID <= 0 {
		return errors.New("Invalid token")
	}

	u.IsVerified = true
	u.DeleteAttribute(user.AttributeConfirmToken)

	if err := c.Crud.Save(&u); err != nil {
		return err
	}

	if err := c.Bus.Emit(context.Background(), user.EventUserConfirm, &u); err != nil {
		return err
	}

	return nil
}

// Restore godoc
func (c *AuthService) Restore(input *RestoreDto) error {
	u, err := c.Crud.FindByIdentity(input.Identity)
	if err != nil {
		return err
	}

	if u.IsDisabled {
		return errors.New("User is disabled")
	}

	u.SetAttribute(user.AttributeResetToken, goutils.RandomToken())

	if err := c.Crud.Save(u); err != nil {
		return err
	}

	if err := c.Bus.Emit(context.Background(), user.EventUserRestore, u); err != nil {
		return err
	}

	return nil
}

// ResetToken godoc
func (c *AuthService) ResetToken(u *user.User, token string) error {
	if u.IsDisabled {
		return errors.New("User is disabled")
	}

	if token != u.GetAttribute(user.AttributeResetToken) {
		return errors.New("Invalid token")
	}

	u.DeleteAttribute(user.AttributeResetToken)
	password := goutils.RandomStr(nil)

	// Password
	pass, err := c.hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = pass

	if err := c.Crud.Save(u); err != nil {
		return err
	}

	var event = user.EventResetToken{
		User:     u,
		Password: password,
	}

	if err := c.Bus.Emit(context.Background(), user.EventUserResetToken, event); err != nil {
		return err
	}

	return nil
}

// NewToken godoc
func (c *AuthService) NewToken(u *user.User) (string, error) {
	claims := NewJwtClaims(u.UUID)

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

// hashPassword godoc
func (c *AuthService) hashPassword(pass string) ([]byte, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return []byte(""), err
	}

	return password, err
}
