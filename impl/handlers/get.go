package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func Get(w http.ResponseWriter, r *http.Request) {
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
		e := errors.NewError(errors.GET_FILMS_JSON_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	films, err := database.GetFilms(query.Limit, query.Page, query.Filter.Genre, query.Filter.Release)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := NewError(errors.GET_FILMS_DB_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := structs.FilmResponce{}
	if films != nil {
		response.Count = len(*films)
		response.Films = *films
		response.Paging.Next = nil
		response.Paging.Prev = nil
		response.Filter = query.Filter
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}

}
