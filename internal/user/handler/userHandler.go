package handlers

import (
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/labstack/echo/v4"
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
	// TODO middleware
	// cooke, _ := c.Cookie("session")
	// if isUserAuthorized(cooke, &api.sessions) {
	// 	login, _ := api.sessions.Load(cooke.Value)
	// 	u, _ := api.users.Load(login)
	//
	// 	user := u.(models.User)
	//
	// 	d := models.LoginData{
	// 		Login:   user.Login,
	// 		Name:    user.Name,
	// 		Surname: user.Surname,
	// 		Email:   user.Email,
	// 	}
	// 	response := models.LoginResponse{
	// 		Status: http.StatusOK,
	// 		Data:   d,
	// 		Msg:    "OK",
	// 	}
	// 	return c.JSON(http.StatusOK, response)
	// }

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
}

func (api *UserHandler) Register(c echo.Context) error {
	newUser := new(models.User)
	err := c.Bind(newUser)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errResp.ErrUnpackingJSON)
	}

	ctx := c.Request().Context()
	response, cookeValue, err := api.UserUsecase.Signup(ctx, newUser)

	if err != nil {
		// TODO send error
	}

	cookie := formCookie(cookeValue)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) Logout(c echo.Context) error {
	// if !isUserAuthorized(cookie, &api.sessions) {
	// 	return c.JSON(http.StatusFailedDependency, errResp.ErrNotLoggedin)
	// }

	cookie, _ := c.Cookie("session")
	ctx := c.Request().Context()
	err := api.UserUsecase.Logout(ctx, cookie.Value)
	if err != nil {
		// TODO error handling
	}

	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	response := models.LogoutResponse{
		Status:     http.StatusOK,
		GoodbyeMsg: "Goodbye, friend!",
	}
	return c.JSON(http.StatusOK, response)
}
