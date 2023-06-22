package nest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

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
		errs := make([]ValidationError, len(ve))
		for i, fe := range ve {
			errs[i] = ValidationError{Field: strings.ToLower(fe.Field()), Msg: fe.Error()}
		}

		return ctx.JSON(http.StatusBadRequest, &echo.Map{"errors": errs})
	}

	return NewHTTPError(http.StatusBadRequest, err.Error())
}
