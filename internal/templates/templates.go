package templates

import (
	"io/fs"
	"net/http"
	"sync"
	"text/template"
)

var (
	templates *template.Template
	once      sync.Once
)

func Init(fs fs.FS) (err error) {
	once.Do(func() {
		templates, err = template.ParseFS(fs, "templates/**/*.html")
	})

	return err
}

func Render(w http.ResponseWriter, name string, data map[string]any) {
	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func RenderDefault(w http.ResponseWriter, data map[string]any) {
	Render(w, "base-layout", data)
}
