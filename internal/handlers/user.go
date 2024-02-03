package handlers

import (
	"fmt"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
