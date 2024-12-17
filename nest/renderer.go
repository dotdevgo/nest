package nest

import "io"

// Renderer tbd
type Renderer interface {
	Render(io.Writer, string, interface{}, Context) error
}
