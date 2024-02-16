package handlers

import (
	"net/http"
	"retrognome/internal/repository"
	"retrognome/internal/template"
)

type RetroHandler struct {
	userRepository  *repository.UserRepository
	retroRepository *repository.RetroRepository
}

func NewRetroHandler(userRepository *repository.UserRepository, retroRepository *repository.RetroRepository) *RetroHandler {
	return &RetroHandler{userRepository: userRepository, retroRepository: retroRepository}
}

func (pageHandler *RetroHandler) LoadHomePage(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "", "head.html", "dashboard.html", "navbar.html")
}

func (pageHandler *RetroHandler) CreateRetro(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Create Retro not implemented"))
	w.WriteHeader(http.StatusOK)
}

func (pageHandler *RetroHandler) CloneRetro(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Clone Retro not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) LoadRetroPage(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Load Retro Page not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) UpdateRetro(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Update Retro not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) DeleteRetro(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Delete Retro not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Create Card not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Delete Card not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Update Card not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) MoveCard(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Move Card not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Create Column not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Delete Column not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) RenameColumn(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Rename Column not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Create Todo not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Delete Todo not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}

func (pageHandler *RetroHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	/// Return message saying not implemented
	w.Write([]byte("Update Todo not implemented"))
	w.WriteHeader(http.StatusNotImplemented)
}
