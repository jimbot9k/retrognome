package main

import (
	"fmt"
	"log"
	"net/http"
	"retrognome/internal/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	fmt.Fprintf(w, "%s: Hi there, I love %s!", r.Method, r.URL.Path[1:])
}

func main() {

	log.Printf("Application Starting")
	configuration, err := utils.LoadApplicationConfiguration()
	if err != nil {
		log.Print("Error loading configuration in config.yaml: ", err)
		return
	}
	log.Printf("Configuration for %s loaded successfully", configuration.AppName)

	/// TODO: Build out a route handler
	http.HandleFunc("/", handler)

	log.Printf("Listening on port %d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil))
}
