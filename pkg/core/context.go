package nest

import (
	"github.com/goava/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Context interface {
	echo.Context
	Resolve(ptr di.Pointer, options ...di.ResolveOption)
}

type context struct {
	echo.Context
	Container *di.Container
}

func (c *context) Resolve(ptr di.Pointer, options ...di.ResolveOption) {
	if err := c.Container.Resolve(ptr, options...); err != nil {
		log.Error(err)
		panic(err)
	}
}
