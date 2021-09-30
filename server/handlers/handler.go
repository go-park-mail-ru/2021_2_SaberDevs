package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"server/server/data"
	"server/server/models"
	"strconv"
	"sync"
	"time"
)

type MyHandler struct {
	sessions sync.Map
	users    sync.Map
}

var feedSize = 5

func NewMyHandler() *MyHandler {
	var handler MyHandler
	for _, user := range data.TestUsers {
		handler.users.Store(user.Login, user)
	}
	return &handler
}

func (api *MyHandler) Login(c echo.Context) error {
	// проверяем активные сессии
	cooke, err := c.Cookie("session")
	if err == nil {
		login, ok := api.sessions.Load(cooke.Value)
		if ok {
			u, _ := api.users.Load(login)
			user := u.(models.User)

			b := models.LoginBody{
				Login:   user.Login,
				Name:    user.Name,
				Surname: user.Surname,
				Email:   user.Email,
			}
			response := models.LoginResponse{
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
	u, ok := api.users.Load(requestUser.Login)
	user := u.(models.User)

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

	// добавляем пользователя в активные сессии
	// api.sessions[cookie.Value] = user.(models.User).Login
	api.sessions.Store(cookie.Value, user.Login)

	// формируем ответ
	b := models.LoginBody{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
		Score:   12345678,
	}
	response := models.LoginResponse{
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

	_, exists := api.users.Load(newUser.Email)

	if exists {
		errorJson := models.ErrorBody{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "User already exists",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	cc, err := c.Cookie("session")
	if err == nil {
		_, exists = api.sessions.Load(cc.Value)

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
	}

	// api.users[newUser.Login] = user
	api.users.Store(newUser.Login, user)

	// ставим куку на сутки
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	//
	// добавляем пользователя в активные сессии
	// api.sessions[cookie.Value] = user.Login
	api.sessions.Store(cookie.Value, user.Login)

	// формируем ответ
	s := models.SignUpBody{
		Login:   user.Login,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}
	response := models.SignupResponse{
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
			GoodbyeMsg: "Goodbuy, friend!",
		}
		return c.JSON(http.StatusOK, response)

	}

	// delete(api.sessions, cookie.Value)
	api.sessions.Delete(cookie.Value)

	// ставим протухшую куку
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	// формируем ответ
	response := models.LogoutResponse{
		Status:     http.StatusOK,
		GoodbyeMsg: "Goodbuy, friend!",
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

		ChunkData = testData[from:to]

	} else {

		start := 0
		if len(testData) > 6 {
			start = len(testData) - 6
		}
		ChunkData = testData[start : len(testData)-1]

	}
	// формируем ответ
	response := models.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}
