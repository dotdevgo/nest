package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func I18n(bundle *i18n.Bundle) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accept := c.Request().Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(bundle, accept)
			c.Set("localizer", localizer)
			return handlerFunc(c)
		}
	}
}
