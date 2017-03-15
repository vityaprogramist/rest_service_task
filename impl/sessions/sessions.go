package sessions

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/rest_service_task/impl/errors"
)

type SessionManager interface {
	SetSession(w http.ResponseWriter, v interface{})
	ReadSession(r *http.Request, v interface{}) error
}

type secureSessionManager struct {
	store *securecookie.SecureCookie
}

func (m *secureSessionManager) SetSession(w http.ResponseWriter, v interface{}) {
	if encoded, err := m.store.Encode("session_id", v); err == nil {
		// cookie := &http.Cookie{
		// 	Name:  "session_id",
		// 	Value: encoded,
		// 	Path:  "/",
		// }
		w.Header().Set("session_id", encoded)
		//http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	} else {
		// log ERROR
		e := errors.NewError(errors.INTERNAL_ERROR)
		if fatal := errors.WriteHttpErrorMessage(w, http.StatusUnauthorized, e); fatal != nil {
			// log FATAL_ERROR
		}
		return
	}
}

func (m *secureSessionManager) ReadSession(r *http.Request, v interface{}) error {
	//cookie, err := r.Cookie("session_id")

	session := r.Header.Get("session_id")

	// if err != nil {
	// 	return err
	// }

	return m.store.Decode("session_id", session, v)
}

func NewSessionManager() SessionManager {
	sc := securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
	return &secureSessionManager{
		store: sc,
	}
}
