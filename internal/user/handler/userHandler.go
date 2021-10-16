package handlers

import (
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/labstack/echo/v4"
	emoji "github.com/tmdvs/Go-Emoji-Utils"

	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
	sessions sync.Map
	users    sync.Map
}
// TODO добавить аргумент юзкейс
func NewUserHandler() *UserHandler {
	var handler UserHandler
	for _, user := range data.TestUsers {
		handler.users.Store(user.Login, user)
	}
	return &handler
}

func formCookie(cookeValue string) *http.Cookie {
	return &http.Cookie{
		Name:     "session",
		Value:    cookeValue,
		HttpOnly: true,
		Expires:  time.Now().Add(10 * time.Hour),
	}
}

func isUserAuthorized(cookie *http.Cookie, sessionsMap *sync.Map) bool {
	if cookie == nil {
		return false
	}
	_, ok := sessionsMap.Load(cookie.Value)
	return ok
}

func (api *UserHandler) Login(c echo.Context) error {
	// TODO middleware
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

	requestUser := new(models.User)
	err := c.Bind(requestUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResp.ErrUnpackingJSON)
	}

	ctx := c.Request().Context()
	response, cookeValue, err := api.UserUsecase.LoginUser(ctx, requestUser)
	if err != nil {
		// TODO send error
	}

	cookie := formCookie(cookeValue)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)

	// u, ok := api.users.Load(requestUser.Login)
	// if !ok {
	// 	return c.JSON(http.StatusFailedDependency, errResp.ErrUserDoesntExist)
	// }
	//
	// user := u.(models.User)
	// if user.Password != requestUser.Password {
	// 	return c.JSON(http.StatusFailedDependency, errResp.ErrWrongPassword)
	// }
	//
	// cookie := formCookie()
	// c.SetCookie(cookie)
	//
	// api.sessions.Store(cookie.Value, user.Login)
	//
	// d := models.LoginData{
	// 	Login:   user.Login,
	// 	Name:    user.Email,
	// 	Surname: user.Email,
	// 	Email:   user.Email,
	// }
	// response := models.LoginResponse{
	// 	Status: http.StatusOK,
	// 	Data:   d,
	// 	Msg:    "OK",
	// }
	//
	// return c.JSON(http.StatusOK, response)
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
	if minLength-emoCount < 0 {
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

func (api *UserHandler) Register(c echo.Context) error {
	newUser := new(models.RequestSignup)
	err := c.Bind(newUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResp.ErrUnpackingJSON)
	}

	_, exists := api.users.Load(newUser.Login)
	if exists {
		return c.JSON(http.StatusFailedDependency, errResp.ErrUserExists)
	}

	cc, _ := c.Cookie("session")
	if isUserAuthorized(cc, &api.sessions) {
		return c.JSON(http.StatusFailedDependency, errResp.ErrAuthorised)
	}

	switch {
	case isValidEmail(newUser.Email):
		return c.JSON(http.StatusFailedDependency, errResp.ErrInvalidEmail)
	case isPasswordValid(newUser.Password):
		return c.JSON(http.StatusFailedDependency, errResp.ErrInvalidPassword)
	case isLoginValid(newUser.Login):
		return c.JSON(http.StatusFailedDependency, errResp.ErrInvalidLogin)
	}

	user := models.User{
		Login:    newUser.Login,
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	api.users.Store(newUser.Login, user)

	cookie := formCookie("")
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

func (api *UserHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session")
	if !isUserAuthorized(cookie, &api.sessions) {
		return c.JSON(http.StatusFailedDependency, errResp.ErrNotLoggedin)
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
