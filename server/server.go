package server

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyHandler struct {
	sessions map[string]uint
	sMu      sync.RWMutex
	users    map[string]*User
	uMu      sync.RWMutex
}

func NewMyHandler() MyHandler {
	return MyHandler{
		sessions: make(map[string]uint, 10),
		users: map[string]*User{
			"mollen@exp.ru": {1, "mollen@exp.ru", "123"},
			"dar@exp.ru": {2, "dar@exp.ru", "123"},
			"viphania@exp.ru": {3, "viphania@exp.ru", "123"},
		},
	}
}

func (api *MyHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func (api *MyHandler) Register(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func (api *MyHandler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func (api *MyHandler) Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func Run() {
	e := echo.New()
	api := NewMyHandler()

	e.POST("/", api.Root)

	e.Logger.Fatal(e.Start(":8080"))
}