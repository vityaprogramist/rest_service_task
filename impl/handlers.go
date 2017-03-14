package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
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
		e := NewError(GET_FILMS_JSON_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	if user.PassInfo == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		e := NewError(REGISTER_PASS_ERROR)
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
		e := NewError(REGISTER_PASS_ERROR)
		e.Message = err.Error()
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	return
}

func Auth(w http.ResponseWriter, r *http.Request) {
	var auth AuthUser
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
		e := NewError(AUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	passHash := HashPassword(auth.Password)

	user, err := database.AuthUser(auth.Login, passHash)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(AUTH_INTERNAL_ERROR)
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
		e := NewError(AUTH_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	var query FilmQuery
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
		e := NewError(GET_FILMS_JSON_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	films, err := database.GetFilms(query.Limit, query.Page, query.Filter.Genre, query.Filter.Release)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := NewError(GET_FILMS_DB_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := FilmResponce{}
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

func GetByID(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")
	if cookie == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(GET_BY_ID_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var user User
	if err := sc.Decode(cookie.Name, cookie.Value, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(GET_BY_ID_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var query FilmQuery
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
		e := NewError(GET_BY_ID_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	films, err := database.GetFilmsByUser(*user.ID, query.Limit, query.Page)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		e := NewError(GET_BY_ID_INTERNAL_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	response := FilmResponce{}
	if films != nil {
		response.Films = *films
		response.Count = len(*films)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func StartRent(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if cookie == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(START_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	filmId, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		e := NewError(START_FILM_ID_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var user User
	if err := sc.Decode(cookie.Name, cookie.Value, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(START_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	err = database.StartRent(user.Login, filmId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		e := ErrorResponse{}
		e.Code = START_ERROR
		e.Message = err.Error()
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func EndRent(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if cookie == nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(END_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	filmId, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		e := NewError(END_FILM_ID_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	var user User
	if err := sc.Decode(cookie.Name, cookie.Value, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		e := NewError(END_UNAUTH_ERROR)
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	err = database.EndRent(filmId, *user.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		e := ErrorResponse{}
		e.Code = END_ERROR
		e.Message = err.Error()
		if err = json.NewEncoder(w).Encode(e); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
