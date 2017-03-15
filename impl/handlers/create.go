package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func (hs *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		e := errors.NewError(errors.GET_FILMS_JSON_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	if user.PassInfo == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		e := errors.NewError(errors.REGISTER_PASS_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	passHash := HashPassword(*user.PassInfo)
	user.PassInfo = &passHash

	err = database.CreateUser(user.FirstName, user.LastName, user.Login, *user.PassInfo, user.Age, user.Phone)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := errors.NewError(errors.REGISTER_PASS_ERROR)
		e.Message = err.Error()
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	return
}
