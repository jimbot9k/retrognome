package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"retrognome/internal/configuration"
	"retrognome/internal/database"
	"retrognome/internal/handlers"
	"retrognome/internal/repository"
	"time"
)

func main() {

	log.Printf("Application Starting")
	configuration, err := configuration.LoadApplicationConfiguration("config.yaml")
	if err != nil {
		log.Print("Error loading configuration in config.yaml: ", err)
		return
	}
	log.Printf("Configuration for %s loaded successfully", configuration.AppName)

	db := database.NewSqliteDB()
	defer db.Close()
	err = database.UpdateSchema(db)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Database loaded successfully")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("web"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	retroRepository := repository.NewRetroRepository(db)

	sessionHandler := handlers.NewSessionHandler(sessionRepository)
	loginHandler := handlers.NewLoginHandler(userRepository, sessionRepository)
	retroHandler := handlers.NewRetroHandler(userRepository, retroRepository)

	router.Get("/register", sessionHandler.RedirectValidSession(loginHandler.LoadRegistrationPage))
	router.Post("/register", sessionHandler.RedirectValidSession(loginHandler.RegisterUser))
	router.Get("/login", sessionHandler.RedirectValidSession(loginHandler.LoadLoginPage))
	router.Post("/login", sessionHandler.RedirectValidSession(loginHandler.LoginUser))
	router.HandleFunc("/logout", sessionHandler.RedirectInvalidSession(loginHandler.DeleteSession))

	router.Get("/", sessionHandler.RedirectInvalidSession(retroHandler.LoadHomePage))
	router.Post("/retro", sessionHandler.RedirectInvalidSession(retroHandler.CreateRetro))
	router.Post("/retro/{retroId}/clone", retroHandler.CloneRetro)
	router.Get("/retro/{retroId}", retroHandler.LoadRetroPage)

	router.Post("/retro/{retroId}/column/{columnId}/card", retroHandler.CreateCard)
	router.Delete("/retro/{retroId}/column/{columnId}/card/{cardId}", retroHandler.DeleteCard)
	router.Put("/retro/{retroId}/column/{columnId}/card/{cardId}", retroHandler.MoveCard)

	router.Post("/retro/{retroId}/column", retroHandler.CreateColumn)
	router.Delete("/retro/{retroId}/column/{columnId}", retroHandler.DeleteColumn)
	router.Put("/retro/{retroId}/column/{columnId}", retroHandler.RenameColumn)

	router.Post("/retro/{retroId}/todo", retroHandler.CreateTodo)
	router.Delete("/retro/{retroId}/todo/{todoId}", retroHandler.DeleteTodo)
	router.Put("/retro/{retroId}/todo/{todoId}", retroHandler.UpdateTodo)

	log.Printf("Listening on http://127.0.0.1:%d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
}
