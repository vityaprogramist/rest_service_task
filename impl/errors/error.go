package errors

import "fmt"

const (
	GET_FILMS_JSON_ERROR int = iota + 1
	GET_FILMS_DB_ERROR
	REGISTER_PASS_ERROR
	AUTH_ERROR
	AUTH_INTERNAL_ERROR
	GET_BY_ID_UNAUTH_ERROR
	GET_BY_ID_INTERNAL_ERROR
	START_UNAUTH_ERROR
	START_FILM_ID_ERROR
	START_ERROR
	END_UNAUTH_ERROR
	END_FILM_ID_ERROR
	END_ERROR
)

const (
	GET_FILMS_JSON_ERROR_MESSAGE     string = "Passed bad json data."
	GET_FILMS_DB_ERROR_MESSAGE       string = "We have some problem with backend. :("
	REGISTER_PASS_ERROR_MESSAGE      string = "Не указан пароль."
	AUTH_ERROR_MESSAGE               string = "Неверный логин или пароль."
	AUTH_INTERNAL_ERROR_MESSAGE      string = "Неизвестная ошибка при авторизации пользователя."
	GET_BY_ID_UNAUTH_ERROR_MESSAGE   string = "Авторизуйтесь!"
	GET_BY_ID_INTERNAL_ERROR_MESSAGE string = "ой, что-то пошло не так. O_o"
	START_UNAUTH_ERROR_MESSAGE       string = "Авторизуйтесь!"
	START_FILM_ID_ERROR_MESSAGE      string = "Нет такого фильма"
)

var messages = map[int]string{
	GET_FILMS_JSON_ERROR:     GET_FILMS_JSON_ERROR_MESSAGE,
	GET_FILMS_DB_ERROR:       GET_FILMS_DB_ERROR_MESSAGE,
	REGISTER_PASS_ERROR:      REGISTER_PASS_ERROR_MESSAGE,
	AUTH_ERROR:               AUTH_ERROR_MESSAGE,
	AUTH_INTERNAL_ERROR:      AUTH_INTERNAL_ERROR_MESSAGE,
	GET_BY_ID_UNAUTH_ERROR:   GET_BY_ID_UNAUTH_ERROR_MESSAGE,
	GET_BY_ID_INTERNAL_ERROR: GET_BY_ID_INTERNAL_ERROR_MESSAGE,
	START_UNAUTH_ERROR:       START_UNAUTH_ERROR_MESSAGE,
	START_FILM_ID_ERROR:      START_FILM_ID_ERROR_MESSAGE,
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
