package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"retrognome/internal/configuration"
	"retrognome/internal/handlers"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("web"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	router.Get("/", handlers.LoadHomePage)
	router.Get("/login", handlers.LoadLoginPage)
	router.Get("/register", handlers.LoadRegistrationPage)
	router.Post("/login", handlers.LoginUser)
	router.Post("/register", handlers.RegisterUser)

	log.Printf("Listening on http://127.0.0.1:%d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
}
