package templates

import (
	"io/fs"
	"net/http"
	"text/template"
)

var templates *template.Template

func Init(fs fs.FS) {
	templates = template.Must(template.ParseFS(fs, "templates/**/*.html"))
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
