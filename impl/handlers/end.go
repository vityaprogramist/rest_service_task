package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rest_service_task/impl/structs"
)

func EndRent(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if cookie == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(END_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	filmId, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		e := NewError(END_FILM_ID_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var user structs.User
	if err := sc.Decode(cookie.Name, cookie.Value, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(END_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	err = database.EndRent(filmId, *user.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		e := ErrorResponse{}
		e.Code = END_ERROR
		e.Message = err.Error()
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
