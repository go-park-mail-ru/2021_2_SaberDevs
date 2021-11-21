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
	ErrorMsg: "User doesnt exist",
}

var ErrWrongPassword = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Wrong password",
}

var ErrUserExists = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "User already exists",
}

var ErrAuthorised = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Already authorised",
}

var ErrNotLoggedin = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Not logged in",
}

var ErrNotFeedNumber = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Not a feed Number",
}

var ErrInvalidEmail = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Invalid symbols in email",
}

var ErrInvalidPassword = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Invalid symbols in password",
}

var ErrInvalidLogin = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Invalid symbols in login",
}

var ErrDbFailure = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Invalid data in DB",
}

var ErrNoSession = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Session doesnt exist",
}

var ErrBadImage = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Cannot upload image",
}

var ErrValidation = ErrorResponse{
	Status: http.StatusNotFound,
}
