package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func (hs *Handlers) Rental(w http.ResponseWriter, r *http.Request) {
	var user structs.User

	err := hs.secure.ReadSession(r, &user)
	if err != nil {
		// log ERROR
		fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, errors.NewError(errors.NOT_AUTHORIZED_ERROR))
		if fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	var query structs.FilmQuery
	err = ReadBody(r, &query)
	if err != nil {
		// log ERROR
		fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
		if fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	films, err := hs.database.GetFilmsByUser(*user.ID, query.Limit, query.Page)
	if err != nil {
		// log ERROR
		e := &errors.ErrorResponse{
			Code:    errors.FORBIDDEN_ERROR,
			Message: err.Error(),
		}

		fatal := errors.WriteHttpErrorMessage(w, http.StatusForbidden, e)
		if fatal != nil {
			// log FATAL_ERROR
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
		// log FATAL_ERROR
	}
}
