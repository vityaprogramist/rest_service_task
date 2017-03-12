package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth", authenticate).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/movies", register).Methods("GET")

	hasher := sha1.New()
	hasher.Write([]byte("vasya"))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	hasher1 := sha1.New()
	hasher1.Write([]byte("vasya"))
	sha1 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	fmt.Printf("hash: %v  ?=   %v", sha, sha1)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
}
