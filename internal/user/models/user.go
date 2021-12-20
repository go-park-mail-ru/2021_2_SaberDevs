package models

import (
	"context"
)

//easyjson:json
type User struct {
	Login       string `json:"login" db:"login"`
	Name        string `json:"firstName" db:"name"`
	Surname     string `json:"lastName" db:"surname"`
	Email       string `json:"email" db:"email" valid:"email,optional" `
	Password    string `json:"password" db:"password"`
	Score       int    `json:"score" db:"score"`
	AvatarURL   string `json:"avatarUrl" db:"avatarurl"`
	Description string `json:"description" db:"description"`
}

//easyjson:json
type RequestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//easyjson:json
type LoginData struct {
	Login       string `json:"login"`
	Surname     string `json:"lastName"`
	Name        string `json:"firstName"`
	Email       string `json:"email"`
	Score       int    `json:"score"`
	AvatarURL   string `json:"avatarUrl"`
	Description string `json:"description"`
}

//easyjson:json
type LoginResponse struct {
	Status uint      `json:"status"`
	Data   LoginData `json:"data"`
	Msg    string    `json:"msg"`
}

// -----------------------------------------------

//easyjson:json
type UpdateProfileData struct {
	Surname     string `json:"lastName"`
	Name        string `json:"firstName"`
	Description string `json:"description"`
}

//easyjson:json
type UpdateProfileResponse struct {
	Status uint              `json:"status"`
	Data   UpdateProfileData `json:"data"`
	Msg    string            `json:"msg"`
}

// -----------------------------------------------

//easyjson:json
type LogoutResponse struct {
	Status     uint   `json:"status"`
	GoodbyeMsg string `json:"goodbye"`
}

// -----------------------------------------------

//easyjson:json
type GetUserData struct {
	Login       string `json:"login"`
	Surname     string `json:"lastName"`
	Name        string `json:"firstName"`
	Score       int    `json:"score"`
	AvatarURL   string `json:"avatarUrl"`
	Description string `json:"description"`
}

//easyjson:json
type GetUserResponse struct {
	Status uint        `json:"status"`
	Data   GetUserData `json:"data"`
	Msg    string      `json:"msg"`
}

// -----------------------------------------------

//easyjson:json
type SignUpData struct {
	Login       string `json:"login"`
	Surname     string `json:"lastName"`
	Name        string `json:"firstName"`
	Email       string `json:"email"`
	Score       int    `json:"score"`
	AvatarURL   string `json:"avatarUrl"`
	Description string `json:"description"`
}

//easyjson:json
type RequestSignup struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"firstName"`
	Surname  string `json:"lastName"`
}

//easyjson:json
type SignupResponse struct {
	Status uint       `json:"status"`
	Data   SignUpData `json:"data"`
	Msg    string     `json:"msg"`
}

// -----------------------------------------------

type UserUsecase interface {
	UpdateProfile(ctx context.Context, user *User, sessionID string) (LoginResponse, error)
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
