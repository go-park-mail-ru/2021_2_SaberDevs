package models

import "context"

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Score    uint   `json:"score"`
}

type RequestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginData struct {
	Login   string `json:"login"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type LoginResponse struct {
	Status uint      `json:"status"`
	Data   LoginData `json:"data"`
	Msg    string    `json:"msg"`
}

// -----------------------------------------------

type LogoutResponse struct {
	Status     uint   `json:"status"`
	GoodbyeMsg string `json:"goodbye"`
}

// -----------------------------------------------

type SignUpData struct {
	Login   string `json:"login"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type RequestSignup struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type SignupResponse struct {
	Status uint       `json:"status"`
	Data   SignUpData `json:"data"`
	Msg    string     `json:"msg"`
}

// -----------------------------------------------

type UserUsecase interface {
	GetByLogin(ctx context.Context, login string) error
	Store(ctx context.Context, user User) error
}

type UserRepository interface {
	GetByLogin(ctx context.Context, login string) error
	Store(ctx context.Context, user User) error
}