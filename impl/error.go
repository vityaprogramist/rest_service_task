package main

import "fmt"

const (
	GET_FILMS_JSON_ERROR int = iota + 1
	GET_FILMS_DB_ERROR
)

const (
	GET_FILMS_JSON_ERROR_MESSAGE string = "Passed bad json data."
	GET_FILMS_DB_ERROR_MESSAGE           string = "We have some problem with backend. :("
)

var messages = map[int]string{
	GET_FILMS_JSON_ERROR: GET_FILMS_JSON_ERROR_MESSAGE,
  GET_FILMS_DB_ERROR : GET_FILMS_DB_ERROR_MESSAGE
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
