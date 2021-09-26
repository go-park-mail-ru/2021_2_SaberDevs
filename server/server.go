package server

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginBody struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	GoodLoginResponse struct {
		Status uint     `json:"status"`
		LBody LoginBody `json:"body"`
	}

	MyHandler struct {
		sessions map[string]uint
		sMu      sync.RWMutex
		users    map[string]User
		uMu      sync.RWMutex
	}
)

func NewMyHandler() MyHandler {
	return MyHandler{
		sessions: make(map[string]uint, 10),
		users: map[string]User{
			"mollen@exp.ru": {1,"mollen@exp.ru", "123"},
			"dar@exp.ru": {2,"dar@exp.ru", "123"},
			"viphania@exp.ru": {3,"viphania@exp.ru", "123"},
		},
	}
}

func (api *MyHandler) Login(c echo.Context) error {
	// достаем данные из запроса
	requestUser := new(User)
	if err := c.Bind(requestUser); err != nil {
		// TODO better error handling
    	return c.String(http.StatusOK, "error reading json")
	}
	// тут что-то про передачу bind полей в функции и небезопасность таких операций ¯\_(ツ)_/¯

	// логика логина
	api.uMu.RLock()
	user, ok := api.users[requestUser.Email]
	api.uMu.RUnlock()

	if !ok {
		// TODO better error handling
		return c.String(http.StatusOK, "no user")
	}
	if user.Password != requestUser.Password {
		// TODO better error handling
		return c.String(http.StatusOK, "wrong password")
	}
	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.ID
	api.sMu.Unlock()

	// формируем ответ
	b := LoginBody{user.ID, user.Email}
	response := GoodLoginResponse{
		Status: http.StatusOK,
		LBody: b,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Register(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func (api *MyHandler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, mollen!")
}

func (api *MyHandler) Root(c echo.Context) error {
	b := LoginBody{11, ""}
	u := GoodLoginResponse{
		Status: 54,
		LBody: b,
	}
	return c.JSON(http.StatusOK, u)
}

func Run() {
	e := echo.New()
	api := NewMyHandler()

	e.POST("api/v1/user/login", api.Login)
	e.GET("/", api.Root)

	e.Logger.Fatal(e.Start(":8080"))
}