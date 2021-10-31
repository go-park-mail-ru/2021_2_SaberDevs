package handlers

import (
	"net/http"
	"sync"
	"time"

	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
}

func NewUserHandler(uu models.UserUsecase) *UserHandler {
	return &UserHandler{uu}
}
func SanitizeUser(a *models.User) *models.User {
	s := bluemonday.StrictPolicy()
	a.Email = s.Sanitize(a.Email)
	a.Login = s.Sanitize(a.Login)
	a.Name = s.Sanitize(a.Name)
	a.Password = s.Sanitize(a.Password)
	//a.Score = s.Sanitize(a.Score)
	a.Surname = s.Sanitize(a.Surname)
	return a
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
	requestUser := new(models.User)
	err := c.Bind(requestUser)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "userHandler/Login",
		}
	}
	requestUser = SanitizeUser(requestUser)
	ctx := c.Request().Context()
	response, sessionID, err := api.UserUsecase.LoginUser(ctx, requestUser)
	if err != nil {
		return errors.Wrap(err, "userHandler/Login")
	}

	cookie := formCookie(sessionID)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) Register(c echo.Context) error {
	newUser := new(models.User)
	err := c.Bind(newUser)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "userHandler.Register",
		}
	}
	newUser = SanitizeUser(newUser)
	ctx := c.Request().Context()
	response, sessionID, err := api.UserUsecase.Signup(ctx, newUser)
	if err != nil {
		return errors.Wrap(err, "userHandler/Register")
	}

	cookie := formCookie(sessionID)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session")
	// TODO middleware

	ctx := c.Request().Context()
	err := api.UserUsecase.Logout(ctx, cookie.Value)
	if err != nil {
		return errors.Wrap(err, "userHandler/Logout")
	}

	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	response := models.LogoutResponse{
		Status:     http.StatusOK,
		GoodbyeMsg: "Goodbye, friend!",
	}
	c.Logger()
	return c.JSON(http.StatusOK, response)
}
