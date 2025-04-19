package nest

import "github.com/defval/di"

// Option defines the signature of a paginator option function.
type Option func(w *Kernel)

// UseContainer godoc
func UseContainer(container *di.Container) Option {
	return func(w *Kernel) {
		w.Container = container
	}
}

// UseProvider godoc
func UseProvider(providers ...di.Option) Option {
	return func(w *Kernel) {
		w.Container.Apply(providers...)
	}
}
