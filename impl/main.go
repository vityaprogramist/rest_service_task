package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var database DBConnection

func main() {
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/auth", authenticate).Methods("POST")
	// router.HandleFunc("/register", register).Methods("POST")
	// router.HandleFunc("/movies", register).Methods("GET")
	//
	// hasher := sha1.New()
	// hasher.Write([]byte("vasya"))
	// sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	//
	// hasher1 := sha1.New()
	// hasher1.Write([]byte("vasya"))
	// sha1 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	//
	// fmt.Printf("hash: %v  ?=   %v", sha, sha1)
	//
	// log.Fatal(http.ListenAndServe(":8080", router))

	f := FilmQuery{}
	g := "Action"
	//r := 2010

	f.Filter.Genre = &g
	f.Filter.Release = nil
	f.Page = 1
	f.Limit = 20

	bs, e := json.Marshal(f)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("json: %s", string(bs))

	database = NewDB()

	err := database.Open("movie_rental", "localhost", 5432, "postgres")
	if err != nil {
		panic(err)
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	return
}

// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
// }
//
// func register(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
// }
