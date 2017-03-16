//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     - api-key:
//       type: apiKey
//       name: session_id
//	     in: header
//
// swagger:meta

package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rest_service_task/impl/db"
	"github.com/rest_service_task/impl/sessions"
)

type Handlers struct {
	database db.DBConnection
	secure   sessions.SessionManager
	logger   *log.Logger
}

func NewHandlerSet(db db.DBConnection, security sessions.SessionManager, log *log.Logger) *Handlers {
	return &Handlers{
		database: db,
		secure:   security,
		logger:   log,
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
