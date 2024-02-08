package handlers

import (
	"net/http"
	"retrognome/internal/repository"
	"retrognome/internal/template"
	"retrognome/internal/types"
)

type PageHandler struct {
	sessionRepository *repository.SessionRepository
}

func NewPageHandler(sessionRepository *repository.SessionRepository) *PageHandler {
	return &PageHandler{sessionRepository: sessionRepository}
}

func (pageHandler *PageHandler) LoadHomePage(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = pageHandler.sessionRepository.GetSessionByToken(token.Value)
	}

	if session.IsEmptySession() {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	template.RenderTemplate(w, "dashboard.html", "")
}

func (pageHandler *PageHandler) LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = pageHandler.sessionRepository.GetSessionByToken(token.Value)
	}

	if session.IsEmptySession() {
		template.RenderTemplate(w, "login.html", "")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (pageHandler *PageHandler) LoadRegistrationPage(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = pageHandler.sessionRepository.GetSessionByToken(token.Value)
	}

	if session.IsEmptySession() {
		template.RenderTemplate(w, "registration.html", "")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
