package injector

import "dotdev/nest/pkg/nest"

// NewSecureGroup godoc
func NewSecureGroup(e *nest.Kernel) nest.SecureGroup {
	g := e.Group("")
	return g
}

// NewApiGroup godoc
// func NewApiGroup(e *nest.Kernel) nest.ApiGroup {
// 	api := e.Group("/api")
// 	return api
// }
