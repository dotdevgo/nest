package nest

import (
	"github.com/go-playground/validator/v10"
	"github.com/goava/di"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	methods = [...]string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		echo.PROPFIND,
		http.MethodPut,
		http.MethodTrace,
		echo.REPORT,
	}
)

// ContainerHandlerFunc TBD
type ContainerHandlerFunc func(Context) interface{}
type HandlerFunc func(Context) error

// EchoWrapper TBD
type EchoWrapper struct {
	*di.Container
	*echo.Echo
}

// New Create new app
func New(m ...di.Option) *EchoWrapper {
	container, _ := di.New(m...)

	e := NewEcho(container)
	w := &EchoWrapper{Container: container, Echo: e}

	if err := container.Provide(func() *EchoWrapper {
		return w
	}); err != nil {
		panic(err)
	}

	return w
}

func NewEcho(container *di.Container) *echo.Echo {
	e := echo.New()

	// Override echo.Context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context{c, container}
			return next(cc)
		}
	})

	// Set custom validator
	e.Validator = &EchoValidator{validator: validator.New()}

	return e
}

// Handler Wrap route with DI args
func (w *EchoWrapper) HandlerFn(handlerFunc ContainerHandlerFunc) HandlerFunc {
	return func(c Context) error {
		//cc := &context{c, w.Container}
		return w.Invoke(handlerFunc(c))
	}
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (w *EchoWrapper) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return w.Echo.Add(method, path, func(c echo.Context) error {
		cc := c.(Context)
		return handler(cc)
	}, middleware...)
}

// CONNECT registers a new CONNECT route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodConnect, path, h, m...)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (w *EchoWrapper) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodDelete, path, h, m...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (w *EchoWrapper) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodGet, path, h, m...)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodHead, path, h, m...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodOptions, path, h, m...)
}

// PATCH registers a new PATCH route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPatch, path, h, m...)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPost, path, h, m...)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPut, path, h, m...)
}

// TRACE registers a new TRACE route for a path with matching handler in the
// router with optional route-level middleware.
func (w *EchoWrapper) TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodTrace, path, h, m...)
}

// Any registers a new route for all HTTP methods and path with matching handler
// in the router with optional route-level middleware.
func (w *EchoWrapper) Any(path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	routes := make([]*echo.Route, len(methods))
	for i, m := range methods {
		routes[i] = w.Add(m, path, handler, middleware...)
	}
	return routes
}

// Match registers a new route for multiple HTTP methods and path with matching
// handler in the router with optional route-level middleware.
func (w *EchoWrapper) Match(methods []string, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	routes := make([]*echo.Route, len(methods))
	for i, m := range methods {
		routes[i] = w.Add(m, path, handler, middleware...)
	}
	return routes
}

// Start starts an HTTP server.
func (w *EchoWrapper) Start(address string) error {
	w.Echo.Server.Addr = address
	return w.Echo.StartServer(w.Echo.Server)
}

//func contextMiddleware(container *di.Container) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			cc := &context{c, container}
//			return next(cc)
//		}
//	}
//}

//func contextMiddleware(next echo.ContainerHandlerFunc) echo.ContainerHandlerFunc {
//	return func(c echo.Context) error {
//		cc := &Context{c}
//		return next(cc)
//	}
//}

// NewEchoWrapper TBD
//func NewEchoWrapper(di *di.Container, e *echo.Echo) *EchoWrapper {
//	w := &EchoWrapper{Container: di, Echo: e}
//	return w
//}

//func (w *EchoWrapper) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
//	return w.Add(http.MethodGet, path, h, m...)
//}




//e := echo.New()
//
//// Override echo.Context
//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cc := &context{c, container}
//		return next(cc)
//	}
//})
//
//// Set custom validator
//e.Validator = &EchoValidator{validator: validator.New()}
//