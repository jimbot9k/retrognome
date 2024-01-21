package main

import (
	"fmt"
	"log"
	"net/http"
	"retrognome/internal/configuration"
	"retrognome/internal/router"
)

func main() {

	log.Printf("Application Starting")
	configuration, err := configuration.LoadApplicationConfiguration()
	if err != nil {
		log.Print("Error loading configuration in config.yaml: ", err)
		return
	}
	log.Printf("Configuration for %s loaded successfully", configuration.AppName)

	router := router.NewRouter()
	router.PathHandler( "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sup Router!"))
	})

	log.Printf("Listening on http://127.0.0.1:%d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
}
