package handlers

import (
	"net/http"
	"retrognome/internal/repository"
	"retrognome/internal/types"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepository}
}

func (userHandler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

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

	if user.Password != password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	/// Set session cookie

	http.Header.Add(w.Header(), "HX-Redirect", "/")
}

func (userHandler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

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

	user := &types.User{Email: email, Password: password}
	err = userHandler.userRepository.CreateUser(user)
	if err != nil {
		http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
		return
	}

	http.Header.Add(w.Header(), "HX-Redirect", "/")
}
