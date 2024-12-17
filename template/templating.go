package template

import (
	"dotdev/nest"
	html "html/template"

	"github.com/defval/di"
)

// New godoc
func New(templates *html.Template) di.Option {
	return di.Options(
		di.Provide(func() *TemplateRenderer {
			return &TemplateRenderer{
				templates: templates,
			}
		}),

		di.Invoke(func(w *nest.Kernel, renderer *TemplateRenderer) {
			w.Renderer = renderer
		}),
	)
}

// w.Echo.Renderer = renderer
// nest.NewExtension(func() *templateExtension {
// 	return &templateExtension{}
// }),
// // parseTemplates takes a directory where html files will reside.
// // It'll check nested dirs and load all files with .html ext.
// func parseTemplates(dir string) (*template.Template, error) {
// 	tmpl := template.New("")
// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if strings.Contains(path, ".html") {
// 			_, err = tmpl.ParseFiles(path)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return tmpl, nil
// }

// tmpls, err := parseTemplates(templatePath)
// if err != nil {
// 	logger.FatalOnError(err)
// }
