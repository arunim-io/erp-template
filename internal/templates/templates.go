package templates

import (
	"io/fs"
	"net/http"
	"sync"
	"text/template"
)

var (
	templatesFS fs.FS
	templates   *template.Template
	once        sync.Once
)

func Init(fs fs.FS) (err error) {
	templatesFS = fs

	once.Do(func() {
		templates, err = template.ParseFS(fs, "templates/layouts/*.html")
	})

	return err
}

func Render(w http.ResponseWriter, name string, data map[string]any) error {
	tpls, err := templates.Clone()
	if err != nil {
		return err
	}

	tpl, err := tpls.ParseFS(templatesFS, "templates/pages/"+name)
	if err != nil {
		return err
	}

	if err := tpl.ExecuteTemplate(w, "base-layout", data); err != nil {
		return err
	}

	return nil
}
