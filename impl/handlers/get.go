package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

type FilmQuery struct {
	genre       string
	releaseyear int
	page        int
	size        int
}

//swagger:parameters getfilms
type GetFilmsParams struct {
	// getting films list
	//
	// unique: true
	// required: true
	// in: query
	// name: string
	genre string
}

// swagger:route GET /get film getfilms
// get films
//		Responses:
//			default: ErrorResponse
//			200:
func (hs *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	var genrePtr *string = nil
	var releaseYearPtr *int = nil
	var page int
	var size int

	if genre := r.URL.Query().Get("g"); genre != "" {
		genrePtr = &genre
	}

	if releaseYear := r.URL.Query().Get("ry"); releaseYear != "" {
		v, err := strconv.Atoi(releaseYear)
		if err != nil {
			// log ERROR
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				// log FATAL_ERROR
			}
			return
		}
		releaseYearPtr = &v
	}

	if pageStr := r.URL.Query().Get("p"); pageStr != "" {
		v, err := strconv.Atoi(pageStr)
		if err != nil {
			// log ERROR
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				// log FATAL_ERROR
			}
			return
		}
		page = v
	}

	if sizeStr := r.URL.Query().Get("sz"); sizeStr != "" {
		v, err := strconv.Atoi(sizeStr)
		if err != nil {
			// log ERROR
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				// log FATAL_ERROR
			}
			return
		}
		size = v
	}

	films, err := hs.database.GetFilms(size, page, genrePtr, releaseYearPtr)
	if err != nil {
		// log ERROR
		fatal := errors.WriteHttpErrorMessage(w, http.StatusInternalServerError, errors.NewError(errors.INTERNAL_ERROR))
		if fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := structs.FilmResponce{}
	if films != nil {
		response.Count = len(*films)
		response.Films = *films
		response.Paging.Next = nil
		response.Paging.Prev = nil
		response.Filter = structs.Filter{Genre: genrePtr, Release: releaseYearPtr}
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		// log FATAL_ERROR
	}
}
