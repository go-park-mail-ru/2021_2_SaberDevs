package models

import "net/http"

type ErrorResponse struct {
	Status   uint   `json:"status"`
	ErrorMsg string `json:"msg"`
}

var ErrUnpackingJSON = ErrorResponse{
	Status:   http.StatusUnprocessableEntity,
	ErrorMsg: "Error unpacking JSON",
}
var ErrUserDoesntExist = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "User doesnt exist",
}

var ErrWrongPassword = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Wrong password",
}

var ErrUserExists = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "User already exists",
}

var ErrAuthorised = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Already authorised",
}

var ErrNotLoggedin = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Not logged in",
}

var ErrNotFeedNumber = ErrorResponse{
	Status:   http.StatusNotFound,
	ErrorMsg: "Not a feed Number",
}

var ErrInvalidEmail = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Invalid symbols in email",
}

var ErrInvalidPassword = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Invalid symbols in password",
}

var ErrInvalidLogin = ErrorResponse{
	Status:   http.StatusFailedDependency,
	ErrorMsg: "Invalid symbols in login",
}
