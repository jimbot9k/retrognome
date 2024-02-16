package handlers

import (
	"net/http"
	"retrognome/internal/repository"
	"retrognome/internal/template"
	"retrognome/internal/types"
	"retrognome/internal/utils"
)

type LoginHandler struct {
	sessionRepository *repository.SessionRepository
	userRepository    *repository.UserRepository
}

func NewLoginHandler(sessionRepository *repository.SessionRepository, userRepository *repository.UserRepository) *LoginHandler {
	return &LoginHandler{sessionRepository: sessionRepository, userRepository: userRepository}
}

func (pageHandler *LoginHandler) LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = pageHandler.sessionRepository.GetSessionByToken(token.Value)
	}

	if session.IsEmptySession() {
		template.RenderTemplate(w, "", "head.html", "login.html")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (pageHandler *LoginHandler) LoadRegistrationPage(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("session_token")
	session := &types.Session{}
	if token != nil {
		session = pageHandler.sessionRepository.GetSessionByToken(token.Value)
	}

	if session.IsEmptySession() {
		template.RenderTemplate(w, "", "head.html", "registration.html")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (userHandler *LoginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusUnauthorized)
		return
	}

	user, err := userHandler.userRepository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}
	if user.IsEmptyUser() {
		http.Error(w, "User does not exist", http.StatusUnauthorized)
		return
	}

	if user.Password != utils.HashPassword(password, user.Salt) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	session := &types.Session{UserID: user.ID, Token: utils.RandomString(128)}
	err = userHandler.sessionRepository.CreateSession(session)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{Name: "session_token", Value: session.Token, Path: "/"}
	http.SetCookie(w, cookie)
	http.Header.Add(w.Header(), "HX-Redirect", "/")
}

func (userHandler *LoginHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	confirmPassword := r.Form.Get("confirmPassword")

	if email == "" || password == "" || confirmPassword == "" {
		http.Error(w, "Email, password, and confirm password are required", http.StatusUnauthorized)
		return
	}

	if password != confirmPassword {
		http.Error(w, "Password and confirm password do not match", http.StatusUnauthorized)
		return
	}

	if len(password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusUnauthorized)
		return
	}

	passwordHasLowercase := false
	passwordHasUppercase := false
	passwordHasNumber := false
	passwordHasSpecial := false
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			passwordHasLowercase = true
		} else if c >= 'A' && c <= 'Z' {
			passwordHasUppercase = true
		} else if c >= '0' && c <= '9' {
			passwordHasNumber = true
		} else {
			passwordHasSpecial = true
		}
	}

	if !passwordHasLowercase || !passwordHasUppercase || !passwordHasNumber || !passwordHasSpecial {
		http.Error(w, "Password must contain at least one lowercase letter, one uppercase letter, one number, and one special character", http.StatusUnauthorized)
		return
	}

	existingUser, err := userHandler.userRepository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}
	if !existingUser.IsEmptyUser() {
		http.Error(w, "User already exists", http.StatusUnauthorized)
		return
	}

	salt := utils.RandomString(128)
	password = utils.HashPassword(password, salt)
	user := &types.User{Email: email, Password: password, Salt: salt}
	err = userHandler.userRepository.CreateUser(user)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}

	http.Header.Add(w.Header(), "HX-Redirect", "/")
}

func (userHandler *LoginHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}
	sessionToken := cookie.Value
	err = userHandler.sessionRepository.DeleteSessionByToken(sessionToken)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}
	cookie = &http.Cookie{Name: "session_token", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, cookie)
	http.Header.Add(w.Header(), "HX-Redirect", "/")
}
