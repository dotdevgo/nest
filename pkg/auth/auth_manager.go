package auth

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"dotdev/nest/pkg/user"
	"dotdev/nest/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/goava/di"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/mustafaturan/bus/v3"
)

const (
	AttributeResetToken   = "reset_token"
	AttributeConfirmToken = "confirmToken"
)

// AuthManager godoc
type AuthManager struct {
	di.Inject
	*bus.Bus
	AuthConfig
	Crud      *user.UserCrud
	Validator echo.Validator
}

// Validate godoc
func (c AuthManager) Validate(input SignInDto) (user.User, error) {
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
func (c AuthManager) NewToken(u user.User) (string, error) {
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
func (c AuthManager) ChangePassword(u user.User, input ChangePasswordDto) error {
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
func (c AuthManager) SignUp(input SignUpDto) (user.User, error) {
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
func (c AuthManager) Confirm(token string) error {
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
func (c AuthManager) Restore(input IdentityDto) error {
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

// OAuth godoc
func (c AuthManager) OAuth(gothUser goth.User) (OAuth, error) {
	oauth := OAuth{}

	result := c.Crud.DB().Preload(clause.Associations).
		Where("unique_id = ? AND provider = ?", gothUser.UserID, gothUser.Provider).
		First(&oauth)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		oauth.Provider = gothUser.Provider
		oauth.UniqueID = gothUser.UserID

		password := utils.RandomStr(nil)
		pass, err := c.hashPassword(password)
		if err != nil {
			return oauth, err
		}

		_, err = mail.ParseAddress(gothUser.Email)
		email := gothUser.Email
		if nil != err {
			// TODO: config domain variable
			email = fmt.Sprintf("%s-%s@4squad.net", oauth.UniqueID, oauth.Provider)
		}

		u := user.User{
			Username:    gothUser.UserID,
			DisplayName: fmt.Sprintf("%s (%s)", gothUser.NickName, gothUser.Name),
			Email:       email,
			Password:    pass,
		}

		// oauthAttribute := fmt.Sprintf("oauth_%s", oauth.Provider)
		// u.SetAttribute(oauthAttribute, oauth.UniqueID)

		oauth.User = &u
		if err := c.Crud.DB().Create(&u).Error; err != nil {
			return oauth, err
		}

		if err := c.Crud.DB().Create(&oauth).Error; err != nil {
			return oauth, err
		}
	}

	return oauth, nil
}

// ResetToken godoc
func (c AuthManager) ResetToken(u user.User, token string) error {
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
func (c AuthManager) Save(u *user.User, input user.UserDto) error {
	if nil != input.RawAttributes {
		u.AddAttributes(input.RawAttributes)
	}

	u.DisplayName = ""
	if len(input.DisplayName) > 0 {
		u.DisplayName = input.DisplayName
	}

	u.Bio = ""
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
		u.SetAttribute(AttributeConfirmToken, utils.RandomToken())
		isEmailChanged = true
	}

	if err := c.Validator.Validate(u); err != nil {
		return err
	}

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	if isEmailChanged {
		if err := c.Bus.Emit(context.Background(), EventUserResetEmail, u); err != nil {
			return err
		}
	}

	return nil
}

// hashPassword godoc
func (c AuthManager) hashPassword(pass string) (password []byte, err error) {
	password, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}

	return password, err
}
