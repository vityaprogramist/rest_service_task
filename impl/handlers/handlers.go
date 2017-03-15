package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rest_service_task/impl/db"
	"github.com/rest_service_task/impl/sessions"
)

type Handlers struct {
	database db.DBConnection
	secure   sessions.SessionManager
}

func NewHandlerSet(db db.DBConnection, security sessions.SessionManager) *Handlers {
	return &Handlers{
		database: db,
		secure:   security,
	}
}

func HashPassword(pass string) string {
	hasher := sha1.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ReadBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Printf("body: %s", string(body))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, v); err != nil {
		return err
	}

	// decoder := json.NewDecoder(r.Body)

	return nil
	//return decoder.Decode(v)
}
