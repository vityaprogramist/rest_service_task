package main

import (
	"github.com/gorilla/mux"
	"github.com/rest_service_task/impl/handlers"
)

func NewRouter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("POST").Path("/create").Name("CreateUser").HandlerFunc(h.Create)
	router.Methods("POST").Path("/auth").Name("Authentificate").HandlerFunc(h.Auth)
	router.Methods("GET").Path("/get").Name("GetFilms").HandlerFunc(h.Get)
	router.Methods("POST").Path("/rental").Name("GetRentedFilms").HandlerFunc(h.Rental)
	router.Methods("POST").Path("/start/{id}").Name("StartRentFilm").HandlerFunc(h.StartRent)
	router.Methods("DELETE").Path("/end/{id}").Name("EndRentFilm").HandlerFunc(h.EndRent)

	// router.Methods("").Path("").Name("").Handler();
	// router.Methods("").Path("").Name("").Handler();

	return router
}

// var routes = Routes{
// 	Route{
// 		"Create",
// 		"POST",
// 		"/create",
// 	},
// 	Route{
// 		"AuthUser",
// 		"POST",
// 		"/auth",
// 		Auth,
// 	},
// 	Route{
// 		"GetFilms",
// 		"POST",
// 		"/get",
// 		Get,
// 	},
// 	Route{
// 		"GetByID",
// 		"POST",
// 		"/rented",
// 		GetByID,
// 	},
// 	Route{
// 		"StartRent",
// 		"PUT",
// 		"/start/{id}",
// 		StartRent,
// 	},
// 	Route{
// 		"EndRent",
// 		"DELETE",
// 		"/end/{id}",
// 		EndRent,
// 	},
//}
