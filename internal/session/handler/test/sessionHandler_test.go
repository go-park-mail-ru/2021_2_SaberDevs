package test

import (
	shandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/handler"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
)

func TestCheckSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	susecase := mock.NewMockSessionUsecase(ctrl)

	handler := shandler.NewSessionHandler(susecase)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		resp := umodels.LoginResponse{}
		susecase.EXPECT().IsSession(gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()

		err := handler.CheckSession(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("fail", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "sesion", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		resp := umodels.LoginResponse{}
		susecase.EXPECT().IsSession(gomock.Any(), gomock.Any()).Return(resp, errors.New("err")).AnyTimes()

		err := handler.CheckSession(c)
		assert.Error(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})
}
