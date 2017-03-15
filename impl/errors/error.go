package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	NOT_AUTHORIZED_ERROR = iota + 1
	BAD_REQUEST_ERROR
	INTERNAL_ERROR
	FORBIDDEN_ERROR
)

const (
	NOT_AUTHORIZED_ERROR_MESSAGE string = "Это действие доступно только авторизованным пользователям."
	BAD_REQUEST_ERROR_MESSAGE    string = "Неправильные параметры запроса"
	INTERNAL_ERROR_MESSAGE       string = "Ой, извините, что-то у нас пошло не так. :("
)

var messages = map[int]string{
	NOT_AUTHORIZED_ERROR: NOT_AUTHORIZED_ERROR_MESSAGE,
	BAD_REQUEST_ERROR:    BAD_REQUEST_ERROR_MESSAGE,
	INTERNAL_ERROR:       INTERNAL_ERROR_MESSAGE,
}

func getMessage(code int) string {
	if _, ok := messages[code]; ok {
		return messages[code]
	}
	return ""
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, getMessage(e.Code))
}

func NewError(code int) *ErrorResponse {
	message := getMessage(code)
	return &ErrorResponse{code, message}
}

func WriteHttpErrorMessage(w http.ResponseWriter, status int, e *ErrorResponse) error {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		return err
	}
	return nil
}
