package nest

import (
	"errors"
	"fmt"
	"net/http"

	"dotdev/nest/logger"

	"github.com/defval/di"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	// "github.com/labstack/gommon/log"
	log "github.com/sirupsen/logrus"
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
	// ContainerHandlerFunc godoc
	ContainerHandlerFunc func(Context) interface{}

	// HandlerFunc godoc
	HandlerFunc func(Context) error

	// Kernel godoc
	Kernel struct {
		*di.Container
		*echo.Echo
		Config
	}

	// Extension godoc
	Extension interface {
		Boot(w *Kernel)
	}

	// Validator is the interface that wraps the Validate function.
	Validator interface {
		Validate(i interface{}) error
	}

	// AbstractController godoc
	AbstractController interface {
		New(w *Kernel)
	}

	// Controller godoc
	Controller struct {
		di.Inject
	}
)

var isBooted = false

// New Create new Nest instance
func New(providers ...di.Option) *Kernel {
	LoadEnv()

	c, err := di.New()
	logger.FatalOnError(err)

	e := NewEcho(c)
	e.HideBanner = true

	w := &Kernel{Container: c, Echo: e}

	/**
	 * TODO: refactor
	 * 	move to initialize* func
	 */

	logger.FatalOnError(c.Provide(func() *Kernel {
		return w
	}))

	// RouteGroup
	logger.FatalOnError(c.Provide(func() *RouteGroup {
		return NewRouteGroup(w)
	}))

	// Config
	w.Config = GetConfig()
	c.Provide(func() Config {
		return w.Config
	})

	c.Apply(providers...)

	return w
}

// NewEcho godoc
func NewEcho(container *di.Container) *echo.Echo {
	e := echo.New()

	// Logger
	logger.Init()
	logger.Logger = log.New()

	e.Logger = logger.GetEchoLogger()
	e.Use(logger.Hook())

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

// NewExtension godoc
func NewExtension(provideFn di.Constructor) di.Option {
	return di.Provide(provideFn, di.As(new(Extension)))
}

// NewController godoc
func NewController(provideFn di.Constructor) di.Option {
	return di.Provide(provideFn, di.As(new(AbstractController)))
}

// InvokeFn Invoke calls the function fn. It parses function parameters. Looks for it in a container.
// And invokes function with them. See Invocation for details.
func (w *Kernel) InvokeFn(invocation di.Invocation, options ...di.InvokeOption) {
	if err := w.Invoke(invocation, options...); err != nil {
		w.Logger.Fatal(err.Error())
	}
}

// ProvideFn Provide provides to container reliable way to build type. The constructor will be invoked lazily on-demand.
// For more information about constructors see Constructor interface. ProvideOption can add additional behavior to
// the process of type resolving.
func (w *Kernel) ProvideFn(constructor di.Constructor, options ...di.ProvideOption) {
	if err := w.Provide(constructor, options...); err != nil {
		w.Logger.Fatal(err.Error())
	}
}

// ResolveFn Resolve resolves type and fills target pointer.
//
//	var server *http.Server
//	if err := container.Resolve(&server); err != nil {
//		// handle error
//	}
func (w *Kernel) ResolveFn(ptr di.Pointer, options ...di.ResolveOption) {
	if err := w.Resolve(ptr, options...); err != nil {
		w.Logger.Fatalf("%s", err)
	}
}

// HandlerFn Wrap route with DI args
func (w *Kernel) HandlerFn(handlerFunc ContainerHandlerFunc) HandlerFunc {
	return func(c Context) error {
		//cc := &context{c, w.Container}
		return w.Invoke(handlerFunc(c))
	}
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (w *Kernel) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return w.Echo.Add(method, path, func(c echo.Context) error {
		cc := c.(Context)
		return handler(cc)
	}, middleware...)
}

// CONNECT registers a new CONNECT route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodConnect, path, h, m...)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (w *Kernel) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodDelete, path, h, m...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (w *Kernel) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodGet, path, h, m...)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodHead, path, h, m...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodOptions, path, h, m...)
}

// PATCH registers a new PATCH route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPatch, path, h, m...)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPost, path, h, m...)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodPut, path, h, m...)
}

// TRACE registers a new TRACE route for a path with matching handler in the
// router with optional route-level middleware.
func (w *Kernel) TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return w.Add(http.MethodTrace, path, h, m...)
}

// Any registers a new route for all HTTP methods and path with matching handler
// in the router with optional route-level middleware.
func (w *Kernel) Any(path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	routes := make([]*echo.Route, len(methods))
	for i, m := range methods {
		routes[i] = w.Add(m, path, handler, middleware...)
	}
	return routes
}

// Match registers a new route for multiple HTTP methods and path with matching
// handler in the router with optional route-level middleware.
func (w *Kernel) Match(methods []string, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	routes := make([]*echo.Route, len(methods))
	for i, m := range methods {
		routes[i] = w.Add(m, path, handler, middleware...)
	}
	return routes
}

// Group creates a new router group with prefix and optional group-level middleware.
func (w *Kernel) Group(prefix string, m ...echo.MiddlewareFunc) (g *Group) {
	g = &Group{prefix: prefix, echo: w}
	g.Use(m...)
	return
}

// WrapHandler wraps `http.Handler` into `echo.HandlerFunc`.
func WrapHandler(h http.Handler) HandlerFunc {
	return func(c Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

// Start starts an HTTP server.
func (w *Kernel) Start(address string) error {
	w.Echo.Server.Addr = address

	return w.Echo.StartServer(w.Echo.Server)
}

// Serve starts an HTTP server on default port.
func (w *Kernel) Serve(address interface{}) error {
	if err := w.Boot(); err != nil {
		return err
	}

	if err := w.Invoke(w.useRouter); err != nil {
		w.Logger.Warn(err.Error())
	}

	var config Config
	w.ResolveFn(&config)

	if address == nil || address == "" {
		address = fmt.Sprintf(":%v", config.HTTP.Port)
	}

	return w.Start(address.(string))
}

// Boot godoc
func (w *Kernel) Boot() error {
	if isBooted {
		return errors.New("boot called twice")
	}

	isBooted = true

	if err := w.Invoke(w.useValidator); err != nil {
		return err
	}

	if err := w.Invoke(w.bootstrap); err != nil {
		w.Logger.Warn(err.Error())
	}

	return nil
}

// bootstrap godoc
func (w *Kernel) bootstrap(providers []Extension) error {
	for _, p := range providers {
		p.Boot(w)
		// if err := p.Boot(w); err != nil {
		// 	return err
		// }
	}

	return nil
}

// useValidator godoc
func (w *Kernel) useValidator() error {
	// Set custom validator
	var v *validator.Validate
	if err := w.Resolve(&v); err == nil {
		w.Validator = &EchoValidator{validator: v}
	}

	if err := w.Provide(func() echo.Validator {
		return w.Validator
	}); err != nil {
		return err
	}

	if err := w.Provide(func() Validator {
		return w.Validator
	}); err != nil {
		return err
	}

	return nil
}

// useRouter godoc
func (w *Kernel) useRouter(controllers []AbstractController) {
	for _, controller := range controllers {
		w.InvokeFn(controller.New)
	}
}