package handlers

import (
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

//swagger:parameters CreateUser
type CreateUserParams struct {
	// Required: true
	// in: body
	Body structs.User
}

// swagger:route POST /create user CreateUser
// Method for creation new user
//		Responses:
//			default: ErrorResponse
//			200:
func (hs *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	err := ReadBody(r, &user)

	if err != nil {
		e := errors.NewError(errors.BAD_REQUEST_ERROR)
		if fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	if user.PassInfo == nil {
		e := errors.NewError(errors.BAD_REQUEST_ERROR)
		if fatal := errors.WriteHttpErrorMessage(w, http.StatusBadRequest, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	passHash := HashPassword(*user.PassInfo)
	user.PassInfo = &passHash

	err = hs.database.CreateUser(user.FirstName, user.LastName, user.Login, *user.PassInfo, user.Age, user.Phone)
	if err != nil {
		e := &errors.ErrorResponse{
			Code:    errors.INTERNAL_ERROR,
			Message: err.Error(),
		}

		if fatal := errors.WriteHttpErrorMessage(w, http.StatusInternalServerError, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
