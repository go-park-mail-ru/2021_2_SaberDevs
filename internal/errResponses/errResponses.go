package models

import "net/http"

type ErrorResponse struct {
	Status   uint   `json:"status"`
	ErrorMsg string `json:"msg"`
}

var ErrInternal = ErrorResponse{
	Status:   http.StatusInternalServerError,
	ErrorMsg: "Internal server error",
}

var ErrUnpackingJSON = ErrorResponse{
	Status:   http.StatusUnprocessableEntity,
	ErrorMsg: "Error unpacking JSON",
}

var ErrUserDoesntExist = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Пользователь не существует",
}

var ErrWrongPassword = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Неверный пароль",
}

var ErrUserExists = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Пользователь уже существует",
}

var ErrAuthorised = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Already authorised",
}

var ErrNotLoggedin = ErrorResponse{
	Status:   http.StatusUnauthorized,
	ErrorMsg: "Not logged in",
}

var ErrNotFeedNumber = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Not a feed Number",
}

var ErrInvalidEmail = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Неверные символы в email",
}

var ErrInvalidPassword = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Неверные символы в пароле",
}

var ErrInvalidLogin = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Неверные символы в логине",
}

var ErrDbFailure = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Invalid data in DB",
}

var ErrNoSession = ErrorResponse{
	Status:   http.StatusUnauthorized,
	ErrorMsg: "Несуществует сессии",
}

var ErrUnauthorized = ErrorResponse{
	Status:   http.StatusUnauthorized,
	ErrorMsg: "Нет прав на данное действие",
}

var ErrBadImage = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Неудалось загрузить изображение. Размер не должен превышать 5мб, поддерживаемые форматы png jpeg",
}

var ErrValidation = ErrorResponse{
	Status: http.StatusNotFound,
}
