package server

import (
	"net/http"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RequestUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LogoutUser struct {
		Email string `json:"email"`
	}

	LoginBody struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	GoodLoginResponse struct {
		Status uint      `json:"status"`
		LBody  LoginBody `json:"body"`
	}

	LogoutResponse struct {
		Status     uint   `json:"status"`
		GoodbuyMsg string `json:"goodbuy"`
	}

	SignUpBody struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	RequestSignup struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GoodSignupResponse struct {
		Status uint       `json:"status"`
		SBody  SignUpBody `json:"body"`
	}

	ErrorBody struct {
		Status   uint   `json:"status"`
		ErrorMsg string `json:"error"`
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
			"mollen@exp.ru":   {1, "mollen@exp.ru", "123"},
			"dar@exp.ru":      {2, "dar@exp.ru", "123"},
			"viphania@exp.ru": {3, "viphania@exp.ru", "123"},
		},
	}
}

func (api *MyHandler) Login(c echo.Context) error {
	// достаем данные из запроса
	requestUser := new(RequestUser)
	if err := c.Bind(requestUser); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Internal server error",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	// тут что-то про передачу bind полей в функции и небезопасность таких операций ¯\_(ツ)_/¯

	// логика логина
	api.uMu.RLock()
	user, ok := api.users[requestUser.Email]
	api.uMu.RUnlock()

	if !ok {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "User doesnt exist",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	if user.Password != requestUser.Password {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Wrong password",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.ID
	api.sMu.Unlock()

	// формируем ответ
	b := LoginBody{user.ID, user.Email}
	response := GoodLoginResponse{
		Status: http.StatusOK,
		LBody:  b,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Register(c echo.Context) error {
	// достаем данные из запроса
	newUser := new(RequestSignup)
	if err := c.Bind(newUser); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Internal server error",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	api.uMu.RLock()
	_, exists := api.users[newUser.Email]
	api.uMu.RUnlock()

	if exists {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "User already exists",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}

	cc, err := c.Cookie("session")

	api.sMu.RLock()
	_, exists = api.sessions[cc.Value]
	api.sMu.RUnlock()

	if err == nil && exists {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "Already authorised",
		}
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	// логика регистрации,  добавляем юзера в мапу
	user := User{uint(len(api.users)), newUser.Email, newUser.Email}
	api.uMu.Lock()
	api.users[newUser.Email] = user
	api.uMu.Unlock()

	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.ID
	api.sMu.Unlock()

	// формируем ответ
	s := SignUpBody{user.ID, user.Email}
	response := GoodSignupResponse{
		Status: http.StatusOK,
		SBody:  s,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Logout(c echo.Context) error {
	// достаем логин из запроса
	logoutUser := new(LogoutUser)
	if err := c.Bind(logoutUser); err != nil {
		errorJson := ErrorBody{
			Status:   http.StatusInternalServerError,
			ErrorMsg: "I fucked up",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, errorJson)
	}
	// удаляем пользователя из активных сессий
	cookie, err := c.Cookie("session")
	if err != nil {
		return err
	}
	api.sMu.Lock()
	delete(api.sessions, cookie.Value)
	api.sMu.Unlock()

	// ставим протухшую куку
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	// формируем ответ
	response := LogoutResponse{
		Status:     http.StatusOK,
		GoodbuyMsg: "Goodbuy, " + logoutUser.Email + "!",
	}
	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Root(c echo.Context) error {
	b := LoginBody{11, ""}
	u := GoodLoginResponse{
		Status: 54,
		LBody:  b,
	}
	return c.JSON(http.StatusOK, u)
}

func Run() {
	e := echo.New()
	api := NewMyHandler()

	e.POST("api/v1/user/login", api.Login)
	e.POST("api/v1/user/register", api.Register)
	e.POST("api/v1/user/logout", api.Logout)
	e.GET("/", api.Root)

	e.Logger.Fatal(e.Start(":8080"))
}
