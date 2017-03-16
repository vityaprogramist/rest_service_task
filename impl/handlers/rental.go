package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

// A RentedQueryFlags contains the query flags for things that list films.
// swagger:parameters rented
type RentedQueryFlags struct {
	// Page number
	Page int `json:"p"`

	// Page size
	Size int `json:"sz"`

	// Session ID key, which have authorized users in header
	//
	// in: header
	// required: true
	SessionID string `json:"session_id"`
}

// swagger:route GET /rented film rented
// Method for ending rent film
//		Responses:
//			default: errorResponse
//			200:
func (hs *Handlers) Rental(w http.ResponseWriter, r *http.Request) {
	var page int
	var size int

	if pageStr := r.URL.Query().Get("p"); pageStr != "" {
		v, err := strconv.Atoi(pageStr)
		if err != nil {
			hs.logger.Println(err.Error())
			fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
			if fatal != nil {
				hs.logger.Fatal(fatal.Error())
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
				hs.logger.Fatal(fatal.Error())
			}
			return
		}
		size = v
	}

	var user structs.User

	err := hs.secure.ReadSession(r, &user)
	if err != nil {
		hs.logger.Println(err.Error())
		fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, errors.NewError(errors.NOT_AUTHORIZED_ERROR))
		if fatal != nil {
			hs.logger.Fatal(fatal.Error())
		}
		return
	}

	films, err := hs.database.GetFilmsByUser(*user.ID, size, page)
	if err != nil {
		hs.logger.Println(err.Error())
		e := &errors.ErrorResponse{
			Code:    errors.FORBIDDEN_ERROR,
			Message: err.Error(),
		}

		fatal := errors.WriteHttpErrorMessage(w, http.StatusForbidden, e)
		if fatal != nil {
			hs.logger.Fatal(fatal.Error())
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

	if fatal := json.NewEncoder(w).Encode(response); fatal != nil {
		hs.logger.Fatal(fatal.Error())
	}
}
