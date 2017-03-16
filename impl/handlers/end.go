package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

// Stop rent a film
// swagger:parameters filmrent
type FilmID struct {
	// ID of the film
	//
	// in: path
	// required: true
	ID int64 `json:"id"`
	// Session ID key, which have authorized users in header
	//
	// in: header
	// required: true
	SessionID string `json:"session_id"`
}

// swagger:route DELETE /end/{id} film filmrent
// Method for ending rent film
//		Responses:
//			default: errorResponse
//			200:
func (hs *Handlers) EndRent(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	err := hs.secure.ReadSession(r, &user)
	if err != nil {
		hs.logger.Println(err.Error())
		e := errors.NewError(errors.NOT_AUTHORIZED_ERROR)
		fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, e)
		if fatal != nil {
			hs.logger.Fatal(fatal.Error())
		}
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		hs.logger.Println(err.Error())
		fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, errors.NewError(errors.BAD_REQUEST_ERROR))
		if fatal != nil {
			hs.logger.Fatal(fatal.Error())
		}
		return
	}

	err = hs.database.EndRent(id, *user.ID)
	if err != nil {
		hs.logger.Println(err.Error())
		e := &errors.ErrorResponse{
			Code:    errors.FORBIDDEN_ERROR,
			Message: err.Error(),
		}

		fatal := errors.WriteHttpErrorMessage(w, http.StatusForbidden, e)
		if fatal != nil {
			hs.logger.Fatal(err.Error())
		}
		return
	}

	hs.secure.SetSession(w, &user)
	w.WriteHeader(http.StatusOK)
}
