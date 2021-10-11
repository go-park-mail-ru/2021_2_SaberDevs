package handlers

import (
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/models"
	"github.com/tmdvs/Go-Emoji-Utils"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type Handler struct {
	sessions sync.Map
	users    sync.Map
}

const chunkSize = 5

func NewHandler() *Handler {
	var handler Handler
	for _, user := range data.TestUsers {
		handler.users.Store(user.Login, user)
	}
	return &handler
}

func formCookie() *http.Cookie {
	return &http.Cookie{
		Name: "session",
		Value: uuid.NewV4().String(),
		HttpOnly: true,
		Expires: time.Now().Add(10 * time.Hour),
	}
}

func isUserAuthorized(cookie *http.Cookie, sessionsMap *sync.Map) bool {
	if cookie == nil {
		return false
	}
	_, ok := sessionsMap.Load(cookie.Value)
	return ok
}

func (api *Handler) Login(c echo.Context) error {
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
	err := c.Bind(requestUser)
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
	return err != nil
}

func isLoginValid(input string) bool {
	var validator *regexp.Regexp
	validator = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$")
	return !validator.MatchString(input)
}

func removeAllAndCount(input string) (string, int) {
	matches := emoji.FindAll(input)
	emoCount := 0

	for _, item := range matches {
		emoCount += item.Occurrences
		emo := item.Match.(emoji.Emoji)
		rs := []rune(emo.Value)
		for _, r := range rs {
			input = strings.ReplaceAll(input, string([]rune{r}), "")
		}
	}

	return input, emoCount
}

func minPasswordLength(emoCount int) int {
	minLength := 8
	if minLength - emoCount < 0 {
		return 0
	}
	return minLength - emoCount
}

func isPasswordValid(input string) bool {
	inputWithoutEmoji, emoCount := removeAllAndCount(input)
	var validator *regexp.Regexp
	minPasswordLength := minPasswordLength(emoCount)
	validator = regexp.MustCompile("^[a-zA-Z0-9[:punct:]]{" + strconv.Itoa(minPasswordLength) + ",20}$")
	return !validator.MatchString(inputWithoutEmoji)
}

func (api *Handler) Register(c echo.Context) error {
	newUser := new(models.RequestSignup)
	err := c.Bind(newUser)
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
	case isPasswordValid(newUser.Password):
		return c.JSON(http.StatusFailedDependency, models.ErrInvalidPassword)
	case isLoginValid(newUser.Login):
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

func (api *Handler) Logout(c echo.Context) error {
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

func (api *Handler) Getfeed(c echo.Context) error {
	rec := c.QueryParam("idLastLoaded")
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
