package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest_service_task/impl/handlers"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	http.HandlerFunc
}

type Routes []Route

func NewRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"Create new user",
			"POST",
			"/create",
			h.EndRent,
		},
		Route{
			"Authetithicate user",
			"POST",
			"/auth",
			h.Auth,
		},
		Route{
			"Get all films",
			"GET",
			"/get",
			h.Get,
		},
		Route{
			"Get films rented by user",
			"GET",
			"/rented",
			h.Rental,
		},
		Route{
			"Start rent a film",
			"POST",
			"/start/{id}",
			h.StartRent,
		},
		Route{
			"End rent a film",
			"DELETE",
			"/end/{id}",
			h.EndRent,
		},
	}

	// router.Methods("POST").Path("/create").Name("CreateUser").HandlerFunc(h.Create)
	// router.Methods("POST").Path("/auth").Name("Authentificate").HandlerFunc(h.Auth)
	// router.Methods("GET").Path("/get").Name("GetFilms").HandlerFunc(h.Get)
	// router.Methods("POST").Path("/rental").Name("GetRentedFilms").HandlerFunc(h.Rental)
	// router.Methods("POST").Path("/start/{id}").Name("StartRentFilm").HandlerFunc(h.StartRent)
	// router.Methods("DELETE").Path("/end/{id}").Name("EndRentFilm").HandlerFunc(h.EndRent)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// router.Methods("").Path("").Name("").Handler();
	// router.Methods("").Path("").Name("").Handler();

	return router
}
