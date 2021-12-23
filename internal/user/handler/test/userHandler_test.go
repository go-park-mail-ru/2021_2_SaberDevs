package test

import (
	"bytes"
	"encoding/json"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	uapp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/user_app"
	"testing"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/user_app/mock"
)

func TestUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("fail", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "sessin", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.UserProfile(c)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.GetUserResponse{}
		app.EXPECT().GetUserProfile(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.UserProfile(c)

		assert.NoError(t, err)
	})
}

func TestAuthorProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user?user=h", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.GetUserResponse{Status: 9, Data: &uapp.GetUserData{Score: 9}}
		app.EXPECT().GetAuthorProfile(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.AuthorProfile(c)

		assert.NoError(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("fail", func(t *testing.T) {
		e := echo.New()

		u := umodels.User{
			Name: "Алексей",
			Surname: "Ffhbcnjd",
			Password: "asfqwr123asd",
		}
		body, _ := json.Marshal(u)

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "sessin", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.UpdateProfile(c)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		u := umodels.User{
			Name: "Алексей",
			Surname: "Ffhbcnjd",
			Password: "asfqwr123asd",
		}
		body, _ := json.Marshal(u)

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.LoginResponse{}
		app.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.UpdateProfile(c)

		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		u := umodels.User{
			Name: "Алексей",
			Surname: "Ffhbcnjd",
			Password: "asfqwr123asd",
		}
		body, _ := json.Marshal(u)

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/login", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.LoginResponse{}
		app.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.Login(c)

		assert.NoError(t, err)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		u := umodels.User{
			Login: "mollen",
			Email: "mollen@mollen.ru",
			Password: "asfqwr123asd",
		}
		body, _ := json.Marshal(u)

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/signup", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.SignupResponse{}
		app.EXPECT().Signup(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.Register(c)

		assert.NoError(t, err)
	})
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock.NewMockUserDeliveryClient(ctrl)
	handler := uhandler.NewUserHandler(app)

	t.Run("fail", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "sessin", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Logout(c)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/user/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		uresp := uapp.Nothing{}
		app.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(&uresp, nil).AnyTimes()

		err := handler.Logout(c)

		assert.NoError(t, err)
	})
}
