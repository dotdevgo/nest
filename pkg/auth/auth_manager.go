package auth

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/user"
	"dotdev/nest/pkg/utils"

	"github.com/defval/di"
	"github.com/dgrijalva/jwt-go"
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
	Validator nest.Validator
}

// Validate godoc
func (c AuthManager) Validate(input SignInDto) (*user.User, error) {
	u, err := c.Crud.LoadUser(input.Identity)
	if err != nil {
		return u, err
	}

	if u.IsDisabled {
		return u, ErrorUserDisabled
	}

	// TODO: util method
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(input.Password)); err != nil {
		return u, ErrorInvalidPassword
	}

	return u, nil
}

// NewToken godoc
func (c AuthManager) NewToken(u *user.User) (string, error) {
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
func (c AuthManager) ChangePassword(u *user.User, input ChangePasswordDto) error {
	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(input.Password)); err != nil {
		return ErrorInvalidPassword
	}

	password, err := hashPassword(input.NewPassword)
	if err != nil {
		return err
	}

	u.Password = password

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	return nil
}

// SignUp godoc
func (c AuthManager) SignUp(input SignUpDto) (*user.User, error) {
	var u *user.User
	utils.Copy(&u, &input)

	u.IsVerified = false
	u.IsDisabled = false

	pass, err := hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	u.Password = pass

	u.SetAttribute(AttributeConfirmToken, utils.RandomToken())

	if err := c.Crud.Flush(u); err != nil {
		return nil, err
	}

	event := EventAuthGeneric{u}
	if err := c.Bus.Emit(context.Background(), EventAuthSignUp, event); err != nil {
		return nil, err
	}

	return u, nil
}

// Confirm godoc
func (c AuthManager) Confirm(token string) error {
	u := c.Crud.FinByConfirmToken(token)
	if u == nil {
		return ErrorUserDisabled
	}

	if u.IsDisabled {
		return ErrorUserDisabled
	}

	u.IsVerified = true
	u.DeleteAttribute(AttributeConfirmToken)

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	event := EventAuthGeneric{u}
	if err := c.Bus.Emit(context.Background(), EventAuthConfirm, event); err != nil {
		return err
	}

	return nil
}

// Restore godoc
func (c AuthManager) Restore(input IdentityDto) error {
	u, err := c.Crud.LoadUser(input.Identity)
	if err != nil {
		return err
	}

	if u.IsDisabled {
		return ErrorUserDisabled
	}

	u.SetAttribute(AttributeResetToken, utils.RandomToken())

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	event := EventAuthGeneric{u}
	if err := c.Bus.Emit(context.Background(), EventAuthRestore, event); err != nil {
		return err
	}

	return nil
}

// ResetToken godoc
func (c AuthManager) ResetToken(u *user.User, token string) error {
	if u.IsDisabled {
		return ErrorUserDisabled
	}

	if token != u.GetAttribute(AttributeResetToken) {
		return ErrorInvalidToken
	}

	u.DeleteAttribute(AttributeResetToken)
	password := utils.RandomStr(nil)

	// Password
	pass, err := hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = pass

	if err := c.Crud.Flush(u); err != nil {
		return err
	}

	event := EventResetToken{u, password}
	if err := c.Bus.Emit(context.Background(), EventAuthResetToken, event); err != nil {
		return err
	}

	return nil
}

// Save godoc
func (c AuthManager) Save(u *user.User, input user.UserDto) error {
	if nil != input.RawAttributes {
		u.SetAttributes(input.RawAttributes)
	}

	u.DisplayName = ""
	if len(input.DisplayName) > 0 {
		u.DisplayName = input.DisplayName
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

	if !isEmailChanged {
		return nil
	}

	event := EventAuthGeneric{u}
	if err := c.Bus.Emit(context.Background(), EventAuthResetEmail, event); err != nil {
		return err
	}

	return nil
}

// OAuth godoc
func (c AuthManager) OAuth(gothUser goth.User) (*OAuth, error) {
	oauth := &OAuth{}

	result := c.Crud.DB().Preload(clause.Associations).
		Where("unique_id = ? AND provider = ?", gothUser.UserID, gothUser.Provider).
		First(&oauth)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return oauth, nil
	}

	oauth.Provider = gothUser.Provider
	oauth.UniqueID = gothUser.UserID

	pass, err := hashPassword(utils.RandomStr(nil))
	if err != nil {
		return oauth, err
	}

	// TODO: config domain variable
	_, err = mail.ParseAddress(gothUser.Email)
	email := gothUser.Email
	if nil != err {
		email = fmt.Sprintf("%s-%s@4squad.net", oauth.UniqueID, oauth.Provider)
	}

	oauth.User = &user.User{
		Username:    gothUser.UserID,
		DisplayName: fmt.Sprintf("%s (%s)", gothUser.NickName, gothUser.Name),
		Email:       email,
		Password:    pass,
	}

	if err := c.Crud.DB().Create(oauth.User).Error; err != nil {
		return oauth, err
	}

	if err := c.Crud.DB().Create(&oauth).Error; err != nil {
		return oauth, err
	}

	return oauth, errors.New("some error occured")
}
