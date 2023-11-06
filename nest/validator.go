package nest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError godoc
type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

// IsValid godoc
func (ctx *context) Validate(input interface{}) error {
	if err := ctx.Bind(input); err != nil {
		return err
	}

	if err := ctx.Context.Validate(input); err != nil {
		return NewValidatorError(ctx, err)
	}

	return nil
}

// EchoValidator godoc
type EchoValidator struct {
	validator *validator.Validate
}

// Validate godoc
func (cv *EchoValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err //echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}

// NewValidatorError godoc
func NewValidatorError(ctx Context, err error) error {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		validationErrors := make([]ValidationError, len(ve))

		for i, fe := range ve {
			validationErrors[i] = ValidationError{Field: strings.ToLower(fe.Field()), Msg: fe.Error()}
		}

		return ctx.JSON(http.StatusBadRequest, &Map{"errors": validationErrors})
	}

	return NewHTTPError(http.StatusBadRequest, err.Error())
}
