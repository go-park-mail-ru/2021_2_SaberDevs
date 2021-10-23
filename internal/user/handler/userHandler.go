package handlers

import (
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"sync"
	"time"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
}

func NewUserHandler(uu models.UserUsecase) *UserHandler {
	return &UserHandler{uu}
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
	return c.JSON(http.StatusOK, response)
}
