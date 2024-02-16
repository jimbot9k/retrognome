package template

import (
	"html/template"

	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, data interface{}, templateNames ...string) {

	location := "internal/template/"
	files := make([]string, 0, len(templateNames))
	for _, templateName := range templateNames {
		files = append(files, filepath.Join(".", location, templateName))
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
