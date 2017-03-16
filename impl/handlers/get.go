package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

// A errorResponse is the default error message that is generated.
//
// swagger:response errorResponse
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 `json:"code"`
		Message error `json:"message"`
	} `json:"body"`
}

// A GetQueryFlags contains the query flags for things that list films.
// swagger:parameters listFilms
type GetQueryFlags struct {
	// Film genre
	Genre string `json:"g"`

	// Genre of film
	ReleaseYear int `json:"ry"`

	// Page number
	Page int `json:"p"`

	// Page size
	Size int `json:"sz"`
}

// swagger:route GET /get film listFilms
// get films
//		Responses:
//			default: errorResponse
//			200:
func (hs *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	var genrePtr *string = nil
	var releaseYearPtr *int = nil
	var page int
	var size int

	if genre := r.URL.Query().Get("g"); genre != "" {
		genre = strings.ToLower(genre)
		genrePtr = &genre
	}

	if releaseYear := r.URL.Query().Get("ry"); releaseYear != "" {
		v, err := strconv.Atoi(releaseYear)
		if err != nil {
			hs.logger.Println(err.Error())
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				hs.logger.Println(err.Error())
			}
			return
		}
		releaseYearPtr = &v
	}

	if pageStr := r.URL.Query().Get("p"); pageStr != "" {
		v, err := strconv.Atoi(pageStr)
		if err != nil {
			hs.logger.Println(err.Error())
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				hs.logger.Println(err.Error())
			}
			return
		}
		page = v
	}

	if sizeStr := r.URL.Query().Get("sz"); sizeStr != "" {
		v, err := strconv.Atoi(sizeStr)
		if err != nil {
			hs.logger.Println(err.Error())
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				hs.logger.Println(err.Error())
			}
			return
		}
		size = v
	}

	if page <= 0 || size <= 0 {
		fatal := errors.WriteHttpErrorMessage(w, http.StatusInternalServerError, errors.NewError(errors.BAD_REQUEST_ERROR))
		if fatal != nil {
			hs.logger.Println(fatal.Error())
		}
		return
	}

	films, err := hs.database.GetFilms(size, page, genrePtr, releaseYearPtr)
	if err != nil {
		hs.logger.Println(err.Error())
		fatal := errors.WriteHttpErrorMessage(w, http.StatusInternalServerError, errors.NewError(errors.INTERNAL_ERROR))
		if fatal != nil {
			hs.logger.Println(fatal.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := structs.FilmResponce{}
	if films != nil {
		response.Count = len(*films)
		response.Films = *films
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		hs.logger.Println(err.Error())
	}
}
