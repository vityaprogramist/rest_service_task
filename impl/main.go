package main

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

var database DBConnection

//var cookieStore *sessions.CookieStore
var sc *securecookie.SecureCookie

func main() {
	database = NewDB()
	err := database.Open("movie_rental", "localhost", 5432, "postgres")
	if err != nil {
		panic(err)
	}

	sc = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	return
}
