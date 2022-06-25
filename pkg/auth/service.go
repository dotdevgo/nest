package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"

	"dotdev/nest/pkg/user"
	"dotdev/nest/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

const (
	AttributeResetToken   = "reset_token"
	AttributeConfirmToken = "confirmToken"
)

// AuthService godoc
type AuthService struct {
	di.Inject
	*bus.Bus
	AuthConfig
	Crud *user.UserCrud
}

// Validate godoc
func (c AuthService) Validate(input SignInDto) (user.User, error) {
	u, err := c.Crud.FindByIdentity(input.Identity)
	if err != nil {
		return u, err
	}

	if u.IsDisabled {
		return u, ErrorUserDisabled
	}

	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(input.Password)); err != nil {
		return u, ErrorInvalidPassword
	}

	return u, nil
}

// NewToken godoc
func (c AuthService) NewToken(u user.User) (string, error) {
	claims := NewJwtClaims(u.ID)

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(c.AuthConfig.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

// ChangePassword godoc
func (c AuthService) ChangePassword(u user.User, input ChangePasswordDto) error {
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(input.Password)); err != nil {
		return ErrorInvalidPassword
	}

	// Password
	pass, err := c.hashPassword(input.NewPassword)
	if err != nil {
		return err
	}
	u.Password = pass

	if err := c.Crud.Flush(&u); err != nil {
		return err
	}

	return nil
}

// SignUp godoc
func (c AuthService) SignUp(input SignUpDto) (user.User, error) {
	var u user.User
	utils.Copy(&u, &input)

	// Defaults
	u.IsVerified = false
	u.IsDisabled = false

	u.SetAttribute(AttributeConfirmToken, utils.RandomToken())

	// Password
	pass, err := c.hashPassword(input.Password)
	if err != nil {
		return u, err
	}
	u.Password = pass

	if err := c.Crud.Flush(&u); err != nil {
		return u, err
	}

	if err := c.Bus.Emit(context.Background(), EventUserSignUp, u); err != nil {
		return u, err
	}

	return u, nil
}

// Confirm godoc
func (c AuthService) Confirm(token string) error {
	var u user.User

	result := c.Crud.Stmt().
		Find(&u, datatypes.JSONQuery("attributes").
			Equals(token, AttributeConfirmToken))

	if result.Error != nil || u.Pk <= 0 {
		return errors.New("Invalid token")
	}

	u.IsVerified = true
	u.DeleteAttribute(AttributeConfirmToken)

	if err := c.Crud.Flush(&u); err != nil {
		return err
	}

	if err := c.Bus.Emit(context.Background(), EventUserConfirm, &u); err != nil {
		return err
	}

	return nil
}

// Restore godoc
func (c AuthService) Restore(input IdentityDto) error {
	u, err := c.Crud.FindByIdentity(input.Identity)
	if err != nil {
		return err
	}

	if u.IsDisabled {
		return ErrorUserDisabled
	}

	u.SetAttribute(AttributeResetToken, utils.RandomToken())

	if err := c.Crud.Flush(&u); err != nil {
		return err
	}

	if err := c.Bus.Emit(context.Background(), EventUserRestore, u); err != nil {
		return err
	}

	return nil
}

// ResetToken godoc
func (c AuthService) ResetToken(u user.User, token string) error {
	if u.IsDisabled {
		return ErrorUserDisabled
	}

	if token != u.GetAttribute(AttributeResetToken) {
		return errors.New("Invalid token")
	}

	u.DeleteAttribute(AttributeResetToken)
	password := utils.RandomStr(nil)

	// Password
	pass, err := c.hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = pass

	if err := c.Crud.Flush(&u); err != nil {
		return err
	}

	var event = EventResetToken{u, password}
	if err := c.Bus.Emit(context.Background(), EventUserResetToken, event); err != nil {
		return err
	}

	return nil
}

// Save godoc
func (c AuthService) Save(u *user.User, input user.UserDto) error {
	if nil != input.RawAttributes {
		u.AddAttributes(input.RawAttributes)
	}

	if len(input.DisplayName) > 0 {
		u.DisplayName = input.DisplayName
	}

	if len(input.Bio) > 0 {
		u.Bio = input.Bio
	}

	if len(input.Username) > 0 && u.Username != input.Username {
		u.Username = input.Username
	}

	isEmailChanged := false
	if len(input.Email) > 0 && u.Email != input.Email {
		u.Email = input.Email
		u.IsVerified = false
		isEmailChanged = true
	}

	if isEmailChanged {
		u.SetAttribute(AttributeConfirmToken, utils.RandomToken())

		if err := c.Bus.Emit(context.Background(), EventUserResetEmail, u); err != nil {
			return err
		}
	}

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	return nil
}

// hashPassword godoc
func (c AuthService) hashPassword(pass string) (password []byte, err error) {
	password, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}

	return password, err
}
