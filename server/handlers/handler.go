package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
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

func formCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = uuid.NewV4().String()
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(10 * time.Hour)
	return cookie
}

func isUserAuthorized(cookie *http.Cookie, sessionsMap *sync.Map) bool {
	if cookie == nil {
		return false
	}
	_, res := sessionsMap.Load(cookie.Value)
	return res
}

func (api *MyHandler) Login(c echo.Context) error {
	cooke, _ := c.Cookie("session")

	if isUserAuthorized(cooke, &api.sessions) {
		login, _ := api.sessions.Load(cooke.Value)
		u, _ := api.users.Load(login)

		user := u.(models.User)

		d := models.LoginData{
			Login:   user.Login,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
		}
		response := models.LoginResponse{
			Status: http.StatusOK,
			Data:   d,
			Msg:    "OK",
		}
		return c.JSON(http.StatusOK, response)
	}

	requestUser := new(models.RequestUser)

	byteContent, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorJson := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "Error unpacking JSON",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	err = json.Unmarshal(byteContent, &requestUser)
	if err != nil {
		errorJson := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "Error unpacking JSON",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	u, ok := api.users.Load(requestUser.Login)
	user := u.(models.User)

	if !ok {
		errorJson := models.ErrorResponse{
			Status:   http.StatusNoContent,
			ErrorMsg: "User doesnt exist",
		}
		return c.JSON(http.StatusNoContent, errorJson)
	}
	if user.Password != requestUser.Password {
		errorJson := models.ErrorResponse{
			Status:   http.StatusForbidden,
			ErrorMsg: "Wrong password",
		}
		return c.JSON(http.StatusForbidden, errorJson)
	}

	cookie := formCookie()
	c.SetCookie(cookie)

	api.sessions.Store(cookie.Value, user.Login)

	d := models.LoginData{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
		Score:   12345678,
	}
	response := models.LoginResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Register(c echo.Context) error {
	newUser := new(models.RequestSignup)
	byteContent, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorJson := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "Error unpacking JSON",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	err = json.Unmarshal(byteContent, &newUser)
	if err != nil {
		errorJson := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "Error unpacking JSON",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	_, exists := api.users.Load(newUser.Login)
	if exists {
		errorJson := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "User already exists",
		}
		return c.JSON(http.StatusFailedDependency, errorJson)
	}

	cc, err := c.Cookie("session")
	if err == nil {
		_, exists = api.sessions.Load(cc.Value)
		if exists {
			errorJson := models.ErrorResponse{
				Status:   http.StatusFailedDependency,
				ErrorMsg: "Already authorised",
			}
			return c.JSON(http.StatusFailedDependency, errorJson)
		}
	}

	user := models.User{
		Login:    newUser.Login,
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	api.users.Store(newUser.Login, user)

	cookie := formCookie()
	c.SetCookie(cookie)

	api.sessions.Store(cookie.Value, user.Login)

	s := models.SignUpData{
		Login:   user.Login,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}
	response := models.SignupResponse{
		Status: http.StatusOK,
		Data:  s,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session")
	if !isUserAuthorized(cookie, &api.sessions) {
		response := models.ErrorResponse{
			Status:   http.StatusFailedDependency,
			ErrorMsg: "Not logged in",
		}
		return c.JSON(http.StatusOK, response)
	}

	api.sessions.Delete(cookie.Value)

	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	response := models.LogoutResponse{
		Status:     http.StatusOK,
		GoodbyeMsg: "Goodbye, friend!",
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
		errorJson := models.ErrorResponse{
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