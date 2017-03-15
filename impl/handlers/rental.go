package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func Rental(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")
	if cookie == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := erros.NewError(errors.GET_BY_ID_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var user structs.User
	if err := sc.Decode(cookie.Name, cookie.Value, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(GET_BY_ID_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var query structs.FilmQuery
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &query); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		e := NewError(GET_BY_ID_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	films, err := database.GetFilmsByUser(*user.ID, query.Limit, query.Page)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := NewError(GET_BY_ID_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	response := structs.FilmResponce{}
	if films != nil {
		response.Films = *films
		response.Count = len(*films)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
