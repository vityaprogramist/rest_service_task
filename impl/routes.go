package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	// router := mux.NewRouter().StrictSlash(true)
	//
	// mux.NewRouter()
	return nil
}
