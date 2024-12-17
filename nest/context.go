package nest

import (
	"bytes"
	"dotdev/logger"
	"errors"
	"net/http"

	"github.com/defval/di"
	"github.com/labstack/echo/v4"
)

var (
	ErrorNotFound = errors.New("not found")
)

type (
	// Context godoc
	Context interface {
		echo.Context

		Resolve(ptr di.Pointer, options ...di.ResolveOption) error
		ResolveFn(ptr di.Pointer, options ...di.ResolveOption)

		NotFound() error

		Validate(input any) error

		// Render renders a template with data and sends a text/html response with status
		// code. Renderer must be registered using `Echo.Renderer`.
		Render(code int, name string, data interface{}) error
	}

	context struct {
		echo.Context
		//echo      *echo.Echo
		Container *di.Container
	}

	Map map[string]interface{}
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
		logger.Panic(err)
	}
}

// NotFound godoc
func (c *context) NotFound() error {
	return c.JSON(http.StatusNotFound, Map{"error": ErrorNotFound.Error()})
}

// Render godoc
func (c *context) Render(code int, name string, data interface{}) (err error) {
	var kernel *Kernel
	if err = c.Container.Resolve(&kernel); err != nil {
		return err
	}

	if kernel.Renderer == nil {
		return echo.ErrRendererNotRegistered
	}

	buf := new(bytes.Buffer)
	if err = kernel.Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}

// Localize translate string
// "github.com/nicksnyder/go-i18n/v2/i18n"
// Localize(msg *i18n.Message) (string, error)
// func (c *context) Localize(msg *i18n.Message) (string, error) {
// 	localizer, ok := c.Get("localizer").(*i18n.Localizer)

// 	if ok {
// 		return localizer.LocalizeMessage(msg)
// 	}

// 	return "", errors.New("cannot find localizer")
// }
