package handlers

import (
	"net/http"
	"retrognome/internal/template"
)

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "index.html", "")
}

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "login.html", "")

}

func LoadRegistrationPage(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "register.html", "")
}
