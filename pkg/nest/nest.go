package nest

import (
	"fmt"
	"net/http"

	"github.com/dotdevgo/nest/pkg/goutils"
	"github.com/go-playground/validator/v10"
	"github.com/goava/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

type (
	// ApiGroup godoc
	ApiGroup interface{}
	// ContainerHandlerFunc godoc
	ContainerHandlerFunc func(Context) interface{}
	// HandlerFunc godoc
	HandlerFunc func(Context) error
	// EchoWrapper godoc
	EchoWrapper struct {
		*di.Container
		*echo.Echo
		Config
	}
	// Controller godoc
	Controller interface {
		Router(e *EchoWrapper)
	}
	// Controller godoc
	ServiceProvider interface {
		Boot(w *EchoWrapper) error
	}
)

// New Create new app
func New(m ...di.Option) *EchoWrapper {
	container, _ := di.New(m...)

	e := NewEcho(container)
	e.HideBanner = true

	w := &EchoWrapper{Container: container, Echo: e}

	goutils.NoErrorOrFatal(container.Provide(func() *EchoWrapper {
		return w
	}))

	return w
}

// NewEcho godoc
func NewEcho(container *di.Container) *echo.Echo {
	e := echo.New()

	// Override echo.Context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &context{ctx, container}
			return next(cc)
		}
	})

	return e
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, message ...interface{}) *echo.HTTPError {
	he := &echo.HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}

	err, ok := he.Message.(error)
	if ok {
		he.Message = err.Error()
	}

	return he
}

// Api godoc
func (w *EchoWrapper) ApiGroup() *Group {
	var api ApiGroup
	w.ResolveFn(&api)
	e := api.(*Group)
	return e
}

// Invoke calls the function fn. It parses function parameters. Looks for it in a container.
// And invokes function with them. See Invocation for details.
func (c *EchoWrapper) InvokeFn(invocation di.Invocation, options ...di.InvokeOption) {
	if err := c.Invoke(invocation, options...); err != nil {
		log.Panic(err)
	}
}

// Provide provides to container reliable way to build type. The constructor will be invoked lazily on-demand.
// For more information about constructors see Constructor interface. ProvideOption can add additional behavior to
// the process of type resolving.
func (c *EchoWrapper) ProvideFn(constructor di.Constructor, options ...di.ProvideOption) {
	if err := c.Provide(constructor, options...); err != nil {
		log.Panic(err)
	}
}

// Resolve resolves type and fills target pointer.
//
//	var server *http.Server
//	if err := container.Resolve(&server); err != nil {
//		// handle error
//	}
func (c *EchoWrapper) ResolveFn(ptr di.Pointer, options ...di.ResolveOption) {
	if err := c.Resolve(ptr, options...); err != nil {
		log.Panic(err)
	}
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

// Group creates a new router group with prefix and optional group-level middleware.
func (e *EchoWrapper) Group(prefix string, m ...echo.MiddlewareFunc) (g *Group) {
	g = &Group{prefix: prefix, echo: e}
	g.Use(m...)
	return
}

// Start starts an HTTP server.
func (w *EchoWrapper) StartWith(address string) error {
	w.beforeStart()

	w.Echo.Server.Addr = address

	return w.Echo.StartServer(w.Echo.Server)
}

// Start starts an HTTP server.
func (w *EchoWrapper) Start() error {
	w.beforeStart()

	var config Config
	w.ResolveFn(&config)

	return w.StartWith(fmt.Sprintf(":%v", config.HTTP.Port))
}

// WrapHandler wraps `http.Handler` into `echo.HandlerFunc`.
func WrapHandler(h http.Handler) HandlerFunc {
	return func(c Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func (w *EchoWrapper) beforeStart() {
	// TODO: refactor
	// Custom
	if err := w.Invoke(w.bootContainer); err != nil {
		log.Error(err)
	}

	if err := w.Invoke(w.registerControllers); err != nil {
		log.Error(err)
	}

	// TODO: refactor
	// Set custom validator
	var v *validator.Validate
	w.ResolveFn(&v)
	w.Validator = &EchoValidator{validator: v}
}

// registerControllers godoc
// TODO: refactor
func (w *EchoWrapper) registerControllers(controllers []Controller) {
	for _, controller := range controllers {
		w.InvokeFn(controller.Router)
	}
}

// bootContainer godoc
// TODO: refactor
func (w *EchoWrapper) bootContainer(providers []ServiceProvider) error {
	for _, p := range providers {
		if err := p.Boot(w); err != nil {
			return err
		}
	}

	return nil
}

// // TODO: move injector
// config, err := GetConfig()
// goutils.NoErrorOrFatal(err)
// container.Provide(func() *Config {
// 	return &config
// })
