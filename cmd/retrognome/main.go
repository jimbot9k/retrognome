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
	userHandler := handlers.NewUserHandler(userRepository, sessionRepository)
	pageHandler := handlers.NewPageHandler(sessionRepository)

	router.Get("/", pageHandler.LoadHomePage)
	router.Get("/register", pageHandler.LoadRegistrationPage)
	router.Get("/login", pageHandler.LoadLoginPage)

	router.Post("/login", userHandler.LoginUser)
	router.Post("/register", userHandler.RegisterUser)
	router.Post("/logout", userHandler.LogoutUser)

	log.Printf("Listening on http://127.0.0.1:%d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
}
