package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	hand "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/handler"
	models "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models/mock"
	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uc1 := mock.NewMockLikesUsecase(ctrl)
	uc2 := mock.NewMockLikesUsecase(ctrl)
	handler := hand.NewLikesHandler(uc1, uc2)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		like := models.LikeData{Ltype: 0, Sign: 1, Id: 1}
		body, _ := json.Marshal(like)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/like", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		uc1.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		uc2.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		err := handler.Rate(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		like := models.LikeData{Ltype: 0, Sign: -1, Id: 1}
		body, _ := json.Marshal(like)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/like", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		uc1.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		uc2.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		err := handler.Rate(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		like := models.LikeData{Ltype: 1, Sign: 1, Id: 1}
		body, _ := json.Marshal(like)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/like", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		uc1.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		uc2.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		err := handler.Rate(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		like := models.LikeData{Ltype: 1, Sign: -1, Id: 1}
		body, _ := json.Marshal(like)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/like", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		uc1.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		uc2.EXPECT().Rating(gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		err := handler.Rate(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}
