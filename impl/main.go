package main

import "fmt"

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

	database := NewDB()
	err := database.Open("movie_rental", "localhost", 5432, "postgres")
	if err != nil {
		fmt.Println(err)
	}

	err = database.CreateUser("aaaaa", "bbbbb", "ccccc", "abcde", 99, 79996450549)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
// }
//
// func register(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
// }
