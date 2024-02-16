package main

import (
	"fmt"
	"log"
	"net/http"
	"retrognome/internal/configuration"
	"retrognome/internal/handlers"
	"retrognome/internal/repository"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	log.Printf("Application Starting")
	configuration, err := configuration.LoadApplicationConfiguration("config.yaml")
	if err != nil {
		log.Print("Error loading configuration in config.yaml: ", err)
		return
	}
	log.Printf("Configuration for %s loaded successfully", configuration.AppName)

	database := repository.NewSqliteDB()
	defer database.Close()
	err = repository.UpdateSchema(database)
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

	userRepository := repository.NewUserRepository(database)
	sessionRepository := repository.NewSessionRepository(database)
	retroRepository := repository.NewRetroRepository(database)

	loginHandler := handlers.NewLoginHandler(sessionRepository, userRepository)
	router.Get("/register", loginHandler.LoadRegistrationPage)
	router.Post("/register", loginHandler.RegisterUser)
	router.Get("/login", loginHandler.LoadLoginPage)
	router.Post("/login", loginHandler.LoginUser)
	router.HandleFunc("/logout", loginHandler.LogoutUser)

	retroHandler := handlers.NewRetroHandler(sessionRepository, userRepository, retroRepository)
	router.Get("/", retroHandler.LoadHomePage)
	router.Post("/retro", retroHandler.CreateRetro)
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
