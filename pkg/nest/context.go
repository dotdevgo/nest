package nest

import (
	"errors"

	"dotdev/nest/pkg/logger"

	"github.com/defval/di"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	// Context godoc
	Context interface {
		echo.Context
		Resolve(ptr di.Pointer, options ...di.ResolveOption) error
		ResolveFn(ptr di.Pointer, options ...di.ResolveOption)
		T(msg *i18n.Message) (string, error)
	}

	context struct {
		echo.Context
		Container *di.Container
	}
)

// Resolve godoc
func (c *context) IsTLS() bool {
	var config Config
	c.ResolveFn(&config)
	return config.HTTP.TLS.Enabled || c.Context.IsTLS()
}

// Resolve godoc
func (c *context) Resolve(ptr di.Pointer, options ...di.ResolveOption) error {
	return c.Container.Resolve(ptr, options...)
}

// ResolveFn godoc
func (c *context) ResolveFn(ptr di.Pointer, options ...di.ResolveOption) {
	if err := c.Container.Resolve(ptr, options...); err != nil {
		// TODO: refactor panic
		logger.Panic(err)
	}
}

// T translate string
func (c *context) T(msg *i18n.Message) (string, error) {
	localizer, ok := c.Get("localizer").(*i18n.Localizer)
	if ok {
		return localizer.LocalizeMessage(msg)
	}
	return "", errors.New("cannot find localizer")
}
