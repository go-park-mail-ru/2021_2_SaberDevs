package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/mail"
	"regexp"
	"server/server/data"
	"server/server/models"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type MyHandler struct {
	sessions  sync.Map
	users     sync.Map
	validator *regexp.Regexp
}

var chunkSize = 5

func NewMyHandler() *MyHandler {
	var handler MyHandler
	handler.validator = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$")
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
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ErrUnpackingJSON)
	}

	err = json.Unmarshal(byteContent, &requestUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ErrUnpackingJSON)
	}

	u, ok := api.users.Load(requestUser.Login)
	if !ok {
		return c.JSON(http.StatusFailedDependency, models.ErrUserDoesntExist)
	}

	user := u.(models.User)
	if user.Password != requestUser.Password {
		return c.JSON(http.StatusFailedDependency, models.ErrWrongPassword)
	}

	cookie := formCookie()
	c.SetCookie(cookie)

	api.sessions.Store(cookie.Value, user.Login)

	d := models.LoginData{
		Login:   user.Login,
		Name:    user.Email,
		Surname: user.Email,
		Email:   user.Email,
	}
	response := models.LoginResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (api *MyHandler) Register(c echo.Context) error {
	newUser := new(models.RequestSignup)
	byteContent, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ErrUnpackingJSON)
	}

	err = json.Unmarshal(byteContent, &newUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ErrUnpackingJSON)
	}

	_, exists := api.users.Load(newUser.Login)
	if exists {
		return c.JSON(http.StatusFailedDependency, models.ErrUserExists)
	}

	cc, _ := c.Cookie("session")
	if isUserAuthorized(cc, &api.sessions) {
		return c.JSON(http.StatusFailedDependency, models.ErrAuthorised)
	}

	switch {
	case isValidEmail(newUser.Email):
		return c.JSON(http.StatusFailedDependency, models.ErrInvalidEmail)
	case api.validator.MatchString(newUser.Password):
		return c.JSON(http.StatusFailedDependency, models.ErrInvalidPassword)
	case api.validator.MatchString(newUser.Login):
		return c.JSON(http.StatusFailedDependency, models.ErrInvalidLogin)
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
		Data:   s,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *MyHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session")
	if !isUserAuthorized(cookie, &api.sessions) {
		return c.JSON(http.StatusFailedDependency, models.ErrNotLoggedin)
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
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, models.ErrNotFeedNumber)
	}
	var ChunkData []models.NewsRecord
	// Возвращаем записи
	testData := data.TestData
	if from >= 0 && from+chunkSize < len(testData) {
		ChunkData = testData[from : from+chunkSize]
	} else {
		start := 0
		if len(testData) > chunkSize {
			start = len(testData) - chunkSize
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
