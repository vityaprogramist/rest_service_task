package main

import (
	"fmt"
	"os"

	"github.com/gorilla/securecookie"
	conf "github.com/rest_service_task/impl/config"
)

var database DBConnection

//var cookieStore *sessions.CookieStore
var sc *securecookie.SecureCookie

func main() {
	config, err := conf.ReadCmd(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)

	// database = NewDB()
	// err := database.Open("movie_rental", "localhost", 5432, "postgres")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// sc = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
	//
	// router := NewRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))

	return
}
