package template

// import (
// 	"dotdev/logger"
// 	"dotdev/nest"
// )

// /* Customer Extension */
// type templateExtension struct {
// 	nest.Extension
// }

// func (templateExtension) Boot(w *nest.Kernel) error {
// 	w.Logger.Info("[Ext] Template: boot")

// 	var renderer *TemplateRenderer
// 	if err := w.Resolve(renderer); err != nil {
// 		logger.Fatal(err)
// 	}

// 	w.Logger.Info("[Ext] Template: set renderer %v", renderer)
// 	// w.Renderer = renderer
// 	w.Echo.Renderer = renderer

// 	return nil
// }
