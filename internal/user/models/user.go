package models

import (
	"context"
)

type User struct {
	Login    string `json:"login" db:"login"`
	Name     string `json:"firstName" db:"name"`
	Surname  string `json:"lastName" db:"surname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Score    int    `json:"score" db:"score"`
}

// -----------------------------------------------

type LoginRequestUser struct {
	Login    string `json:"login" valid:"login"`
	Password string `json:"password" valid:"pass"`
}

type LoginData struct {
	Login   string `json:"login"`
	Surname string `json:"lastName"`
	Name    string `json:"firstName"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type LoginResponse struct {
	Status uint      `json:"status"`
	Data   LoginData `json:"data"`
	Msg    string    `json:"msg"`
}

// -----------------------------------------------

type UpdateRequestUser struct {
	Surname  string `json:"lastName"`
	Name     string `json:"firstName"`
	Password string `json:"password" valid:"pass"`
}

type UpdateProfileData struct {
	Surname string `json:"lastName"`
	Name    string `json:"firstName"`
}

type UpdateProfileResponse struct {
	Status uint              `json:"status"`
	Data   UpdateProfileData `json:"data"`
	Msg    string            `json:"msg"`
}

// -----------------------------------------------

type LogoutResponse struct {
	Status     uint   `json:"status"`
	GoodbyeMsg string `json:"goodbye"`
}

// -----------------------------------------------

type GetUserData struct {
	Login   string `json:"login"`
	Surname string `json:"lastName"`
	Name    string `json:"firstName"`
	Score   int    `json:"score"`
}

type GetUserResponse struct {
	Status uint        `json:"status"`
	Data   GetUserData `json:"data"`
	Msg    string      `json:"msg"`
}

// -----------------------------------------------

type SignupRequestUser struct {
	Login    string `json:"login" valid:"login"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"pass"`
}

type SignUpData struct {
	Login   string `json:"login"`
	Surname string `json:"lastName"`
	Name    string `json:"firstName"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type SignupResponse struct {
	Status uint       `json:"status"`
	Data   SignUpData `json:"data"`
	Msg    string     `json:"msg"`
}

// -----------------------------------------------

type UserUsecase interface {
	UpdateProfile(ctx context.Context, user *User, sessionID string) (UpdateProfileResponse, error)
	GetAuthorProfile(ctx context.Context, author string) (GetUserResponse, error)
	GetUserProfile(ctx context.Context, sessionID string) (GetUserResponse, error)
	LoginUser(ctx context.Context, user *User) (LoginResponse, string, error)
	Signup(ctx context.Context, user *User) (SignupResponse, string, error)
	Logout(ctx context.Context, cookieValue string) error
}

type UserRepository interface {
	UpdateUser(ctx context.Context, user *User) (User, error)
	GetByLogin(ctx context.Context, login string) (User, error)
	GetByName(ctx context.Context, name string) (User, error)
	Store(ctx context.Context, user *User) (User, error)
}
