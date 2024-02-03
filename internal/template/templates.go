package template

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {

	location := "internal/template/"
	tmpl, err := template.ParseFiles(filepath.Join(".", location+"head.html"), filepath.Join(".", location+templateName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
