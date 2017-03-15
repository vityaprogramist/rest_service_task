package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rest_service_task/impl/errors"
	"github.com/rest_service_task/impl/structs"
)

func (hs *Handlers) Auth(w http.ResponseWriter, r *http.Request) {
	var auth structs.AuthUser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &auth); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		e := errors.NewError(errors.AUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	passHash := HashPassword(auth.Password)

	user, err := hs.database.AuthUser(auth.Login, passHash)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := errors.NewError(errors.AUTH_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	if encoded, err := sc.Encode("session", user); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := errros.NewError(errors.AUTH_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}
}
