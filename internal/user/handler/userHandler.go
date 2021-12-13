package handlers

import (
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberValidation"
	"net/http"
	"sync"
	"time"

	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/user_app"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

type UserHandler struct {
	UserUsecase app.UserDeliveryClient
}

func NewUserHandler(uu app.UserDeliveryClient) *UserHandler {
	return &UserHandler{uu}
}
func SanitizeUser(a *app.User) *app.User {
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
		Path:     "/",
	}
}

func isUserAuthorized(cookie *http.Cookie, sessionsMap *sync.Map) bool {
	if cookie == nil {
		return false
	}
	_, ok := sessionsMap.Load(cookie.Value)
	return ok
}

func (api *UserHandler) UserProfile(c echo.Context) error {
	sessionID, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNotLoggedin{
			Reason:   err.Error(),
			Function: "userUsecase/UpdateProfile",
		}
	}

	ctx := c.Request().Context()
	grpcSessionID := &app.SessionID{
		SesionID:             sessionID.Value,
	}
	response, err := api.UserUsecase.GetUserProfile(ctx, grpcSessionID)
	if err != nil {
		return errors.Wrap(err, "userHandler/UserProfile")
	}

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) AuthorProfile(c echo.Context) error {
	authorName := c.QueryParam("user")
	ctx := c.Request().Context()
	grpcAuthorName := &app.Author{Author: authorName}

	grpcResponse, err := api.UserUsecase.GetAuthorProfile(ctx, grpcAuthorName)
	if err != nil {
		return errors.Wrap(err, "userHandler/AuthorProfile")
	}

	response := models.LoginResponse{
		Status: uint(grpcResponse.Status),
		Data:   models.LoginData{
			Login:       grpcResponse.Data.Login,
			Name:        grpcResponse.Data.Name,
			Surname:     grpcResponse.Data.Surname,
			Email:       grpcResponse.Data.Email,
			Score:       int(grpcResponse.Data.Score),
			AvatarURL:   grpcResponse.Data.AvatarUrl,
			Description: grpcResponse.Data.Description,
		},
		Msg:    grpcResponse.Msg,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) UpdateProfile(c echo.Context) error {
	requestUser := new(models.User)
	err := c.Bind(requestUser)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "userHandler/Login",
		}
	}

	sessionID, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNotLoggedin{
			Reason:   err.Error(),
			Function: "userUsecase/UpdateProfile",
		}
	}

	err = syberValidation.ValidateUpdate(*requestUser)
	if err != nil {
		return sbErr.ErrValidate{
			Reason:   err.Error(),
			Function: "userHandler/UpdateProfile",
		}
	}

	grpcUpdateInput := &app.UpdateInput{
		User:                 &app.User{
			Name:                 requestUser.Name,
			Surname:              requestUser.Surname,
			Password:             requestUser.Password,
			AvatarUrl:            requestUser.AvatarURL,
			Description:          requestUser.Description,
		},
		SesionID:             sessionID.Value,
	}

	ctx := c.Request().Context()
	response, err := api.UserUsecase.UpdateProfile(ctx, grpcUpdateInput)
	if err != nil {
		return errors.Wrap(err, "userHandler/UpdateProfile")
	}

	return c.JSON(http.StatusOK, response)
}

func (api *UserHandler) Login(c echo.Context) error {
	requestUser := new(app.User)
	err := c.Bind(requestUser)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "userHandler/Login",
		}
	}
	requestUser = SanitizeUser(requestUser)
	ctx := c.Request().Context()
	grpcResponse, err := api.UserUsecase.LoginUser(ctx, requestUser)
	if err != nil {
		return errors.Wrap(err, "userHandler/Login")
	}

	cookie := formCookie(grpcResponse.SessionID)
	c.SetCookie(cookie)
	grpcResponse.SessionID = ""

	return c.JSON(http.StatusOK, grpcResponse)
}

func (api *UserHandler) Register(c echo.Context) error {
	newUser := new(app.User)
	err := c.Bind(newUser)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "userHandler.Register",
		}
	}

	err = syberValidation.ValidateSignUp(*newUser)
	if err != nil {
		return sbErr.ErrValidate{
			Reason:   err.Error(),
			Function: "userHandler/register",
		}
	}

	newUser = SanitizeUser(newUser)
	ctx := c.Request().Context()
	grpcResponse, err := api.UserUsecase.Signup(ctx, newUser)
	if err != nil {
		return errors.Wrap(err, "userHandler/Register")
	}

	cookie := formCookie(grpcResponse.SessionID)
	c.SetCookie(cookie)
	grpcResponse.SessionID = ""

	return c.JSON(http.StatusOK, grpcResponse)
}

func (api *UserHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session")
	// TODO middleware

	ctx := c.Request().Context()
	grpcSessionID := &app.CookieValue{
		CookieValue:             cookie.Value,
	}

	_, err := api.UserUsecase.Logout(ctx, grpcSessionID)
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
