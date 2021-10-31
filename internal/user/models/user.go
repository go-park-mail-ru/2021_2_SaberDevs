package models

import (
	"context"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

type User struct {
	Login    string `json:"login" db:"login"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email" valid:"email,optional" `
	Password string `json:"password" db:"password"`
	Score    int    `json:"score" db:"score"`
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

type UpdateProfileData struct {
	Surname string `json:"surname"`
	Name    string `json:"name"`
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
	Login    string            `json:"login"`
	Surname  string            `json:"surname"`
	Name     string            `json:"name"`
	Score    int               `json:"score"`
	Articles []amodels.Article `json:"articles"`
}

type GetUserResponse struct {
	Status uint        `json:"status"`
	Data   GetUserData `json:"data"`
	Msg    string      `json:"msg"`
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
