package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	conf "github.com/rest_service_task/impl/config"
	"github.com/rest_service_task/impl/db"
	"github.com/rest_service_task/impl/handlers"
	"github.com/rest_service_task/impl/sessions"
	"github.com/rest_service_task/impl/structs"
)

func main() {
	config, err := conf.ReadCmd(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	database := db.NewDB()
	err = database.Open(config.Database, config.DBHost, config.DBPort, config.User, config.Password)
	if err != nil {
		log.Fatalf("FATAL: %s\n", err.Error())
	}

	gob.Register(structs.User{})
	s := sessions.NewSessionManager()
	h := handlers.NewHandlerSet(database, s)
	router := NewRouter(h)

	docHandler := http.StripPrefix("/doc/", http.FileServer(http.Dir("../doc/")))
	router.PathPrefix("/doc/").Handler(docHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", nil))

	return
}
