package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func (hs *Handlers) EndRent(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	err := hs.secure.ReadSession(r, &user)
	if err != nil {
		e := errors.NewError(errors.NOT_AUTHORIZED_ERROR)
		fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, e)
		if fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		// log ERROR
		fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
		if fatal != nil {
			// FATAL_ERROR
		}
		return
	}

	err = hs.database.EndRent(id, *user.ID)
	if err != nil {
		e := &errors.ErrorResponse{
			Code:    errors.FORBIDDEN_ERROR,
			Message: err.Error(),
		}

		fatal := errors.WriteHttpErrorMessage(w, http.StatusForbidden, e)
		if fatal != nil {
			// log  FATAL_ERROR
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
