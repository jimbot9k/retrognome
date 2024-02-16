package handlers

import (
	"log"
	"net/http"
	"retrognome/internal/repository"
	"retrognome/internal/types"
)

type SessionHandler struct {
	sessionRepository *repository.SessionRepository
}

func NewSessionHandler(sessionRepository *repository.SessionRepository) *SessionHandler {
	return &SessionHandler{sessionRepository: sessionRepository}
}

func (sessionHandler *SessionHandler) CheckSession(r *http.Request) *types.Session {
	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = sessionHandler.sessionRepository.GetSessionByToken(token.Value)
	}
	return session
}

func (sessionHandler *SessionHandler) RedirectInvalidSession(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sessionHandler.CheckSession(r).IsEmptySession() {
			log.Print("Session is empty")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (sessionHandler *SessionHandler) RedirectValidSession(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !sessionHandler.CheckSession(r).IsEmptySession() {

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
