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
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/register",
		Register,
	},
	Route{
		"AuthUser",
		"POST",
		"/auth",
		Auth,
	},
	Route{
		"GetFilms",
		"POST",
		"/get",
		Get,
	},
	Route{
		"GetByID",
		"POST",
		"/get/{id}",
		GetByID,
	},
	Route{
		"StartRent",
		"POST",
		"/start",
		StartRent,
	},
	Route{
		"EndRent",
		"POST",
		"/end",
		EndRent,
	},
}
