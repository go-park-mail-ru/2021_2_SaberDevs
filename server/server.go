package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/server/data"
	"strconv"
	"sync"
	"time"

	"server/server/models"

	"github.com/labstack/echo/v4/middleware"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
)

type (
	MyHandler struct {
		sessions map[string]string
		sMu      sync.RWMutex
		users    map[string]models.User
		uMu      sync.RWMutex
	}
)

var feedSize = 5

func NewMyHandler() MyHandler {
	return MyHandler{
		sessions: make(map[string]string, 10),
		users:    data.TestUsers,
	}
}

func (api *MyHandler) Login(c echo.Context) error {
	// проверяем активные сессии
	cooke, err := c.Cookie("session")
	if err == nil {
		api.sMu.RLock()
		login, ok := api.sessions[cooke.Value]
		api.sMu.RUnlock()
		if ok {
			api.uMu.RLock()
			user, _ := api.users[login]
			api.uMu.RUnlock()

			b := models.LoginBody{
				Login:   user.Login,
				Name:    user.Email,
				Surname: user.Email,
				Email:   user.Email,
				Score:   12345678,
			}
			response := models.GoodLoginResponse{
				Status: http.StatusOK,
				Data:   b,
				Msg:    "OK",
			}

			return c.JSON(http.StatusOK, response)
		}
	}
	// достаем данные из запроса
	// requestUser := new(RequestUser)
	// if err := c.Bind(requestUser); err != nil {
	//	errorJson := ErrorBody{
	//		Status:   http.StatusBadRequest,
	//		ErrorMsg: "Json request in wrong format",
	//	}
	//	c.Logger().Printf("Error: %s", err.Error())
	//	return c.JSON(http.StatusBadRequest, errorJson)
	// }
	requestUser := new(models.RequestUser)
	byteContent, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(byteContent, &requestUser)
	if err != nil {
		log.Println(err)
	}
	c.Logger().Printf("login")
	// тут что-то про передачу bind полей в функции и небезопасность таких операций ¯\_(ツ)_/¯

	// логика логина
	api.uMu.RLock()
	user, ok := api.users[requestUser.Login]
	api.uMu.RUnlock()

	if !ok {
		errorJson := models.ErrorBody{
			Status:   http.StatusNoContent,
			ErrorMsg: "User doesnt exist",
		}
		return c.JSON(http.StatusNoContent, errorJson)
	}
	if user.Password != requestUser.Password {
		errorJson := models.ErrorBody{
			Status:   http.StatusForbidden,
			ErrorMsg: "Wrong password",
		}
		return c.JSON(http.StatusForbidden, errorJson)
	}
	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(10 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.Login
	api.sMu.Unlock()

	// формируем ответ
	b := models.LoginBody{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
		Score:   12345678,
	}
	response := models.GoodLoginResponse{
		Status: http.StatusOK,
		Data:   b,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Register(c echo.Context) error {
	// достаем данные из запроса
	// newUser := new(RequestSignup)
	// if err := c.Bind(newUser); err != nil {
	//	errorJson := ErrorBody{
	//		Status:   http.StatusBadRequest,
	//		ErrorMsg: "Json request in wrong format",
	//	}
	//	c.Logger().Printf("Error: %s", err.Error())
	//	return c.JSON(http.StatusBadRequest, errorJson)
	// }
	newUser := new(models.RequestSignup)
	byteContent, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(byteContent, &newUser)
	if err != nil {
		log.Println(err)
	}
	api.uMu.RLock()
	_, exists := api.users[newUser.Email]
	api.uMu.RUnlock()

	if exists {
		errorJson := models.ErrorBody{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "User already exists",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	cc, err := c.Cookie("session")
	if err == nil {
		api.sMu.RLock()
		_, exists = api.sessions[cc.Value]
		api.sMu.RUnlock()

		if exists {
			errorJson := models.ErrorBody{
				Status:   http.StatusFailedDependency,
				ErrorMsg: "Already authorised",
			}
			return c.JSON(http.StatusFailedDependency, errorJson)
		}
	}
	// логика регистрации, добавляем юзера в мапу
	user := models.User{
		Login:    newUser.Login,
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Email:    newUser.Email,
		Password: newUser.Password,
		Score:    12345,
	}
	api.uMu.Lock()
	api.users[newUser.Login] = user
	api.uMu.Unlock()

	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	api.sMu.Lock()
	api.sessions[cookie.Value] = user.Login
	api.sMu.Unlock()

	// формируем ответ
	s := models.SignUpBody{
		Login:   user.Login,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Score:   12345678, // rand.Int(),
	}
	response := models.GoodSignupResponse{
		Status: http.StatusOK,
		SBody:  s,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Logout(c echo.Context) error {
	// удаляем пользователя из активных сессий
	cookie, err := c.Cookie("session")
	if err != nil {
		response := models.LogoutResponse{
			Status:     http.StatusOK,
			GoodbuyMsg: "Goodbuy, friend!",
		}
		return c.JSON(http.StatusOK, response)

	}
	api.sMu.Lock()
	delete(api.sessions, cookie.Value)
	api.sMu.Unlock()

	// ставим протухшую куку
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	// формируем ответ
	response := models.LogoutResponse{
		Status:     http.StatusOK,
		GoodbuyMsg: "Goodbuy, friend!",
	}
	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Getfeed(c echo.Context) error {
	rec := c.QueryParam("idLastLoaded")
	// TODO костыль!!!!
	if rec == "" {
		rec = "0"
	}

	from, err := strconv.Atoi(rec)
	if err != nil {
		errorJson := models.ErrorBody{
			Status:   http.StatusNotFound,
			ErrorMsg: "Not a feed Number",
		}
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, errorJson)
	}
	to := from + 4
	var ChunkData []models.NewsRecord
	// Возвращаем записи
	testData := data.TestData
	if from >= 0 && to < len(testData) {
		api.sMu.RLock()
		ChunkData = testData[from:to]
		api.sMu.RUnlock()
	} else {
		api.sMu.RLock()
		start := 0
		if len(testData) > 6 {
			start = len(testData) - 6
		}
		ChunkData = testData[start : len(testData)-1]
		api.sMu.RUnlock()
	}
	// формируем ответ
	response := models.ChunkResponse{
		Status: http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://87.228.2.178:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))
	api := NewMyHandler()

	e.POST("/login", api.Login)
	e.POST("/signup", api.Register)
	e.POST("/logout", api.Logout)
	e.GET("/feed", api.Getfeed)

	e.Logger.Fatal(e.Start(address))
}
