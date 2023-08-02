package test

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"

	capp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app/mock"
	chandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/handler"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	"testing"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockCommentDeliveryClient(ctrl)

	handler := chandler.NewCommentHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		comment := cmodels.Comment{}
		body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments/create", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentResponse{Data: &capp.PreparedComment{Id: 1, Author: &capp.Author{Login: "mollen"}}}
		app.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.CreateComment(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("fail", func(t *testing.T) {
		e := echo.New()
		comment := cmodels.Comment{}
		body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments/create", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "sessio", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentResponse{Data: &capp.PreparedComment{Id: 1, Author: &capp.Author{Login: "mollen"}}}
		app.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.CreateComment(c)
		assert.Error(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockCommentDeliveryClient(ctrl)

	handler := chandler.NewCommentHandler(app)
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		comment := cmodels.Comment{}
		body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments/update", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentResponse{Data: &capp.PreparedComment{Id: 1, Author: &capp.Author{Login: "mollen"}}}
		app.EXPECT().UpdateComment(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.UpdateComment(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("fail", func(t *testing.T) {
		e := echo.New()
		comment := cmodels.Comment{}
		body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments/update", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		// cc := &http.Cookie{Name: "session", Value: "123"}
		// req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentResponse{Data: &capp.PreparedComment{Id: 1, Author: &capp.Author{Login: "mollen"}}}
		er := errors.New("err")
		app.EXPECT().UpdateComment(gomock.Any(), gomock.Any()).Return(&new, er).AnyTimes()
		err := handler.UpdateComment(c)
		assert.Error(t, err)
	})
}

func TestGetCommentsByArticleID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockCommentDeliveryClient(ctrl)

	handler := chandler.NewCommentHandler(app)

	t.Run("fail", func(t *testing.T) {
		e := echo.New()
		comment := cmodels.Comment{}
		body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments/create", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentChunkResponse{Data: []*capp.PreparedComment{}}
		new.Data = append(new.Data, &capp.PreparedComment{Id: 1, Author: &capp.Author{Login: "mollen"}})

		app.EXPECT().GetCommentsByArticleID(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetCommentsByArticleID(c)
		assert.Error(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		// comment := cmodels.Comment{}
		// body, _ := json.Marshal(comment)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/comments?id=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := capp.CommentChunkResponse{Data: []*capp.PreparedComment{}}
		new.Data = append(new.Data, &capp.PreparedComment{Id: 1, ArticleId: 1, Author: &capp.Author{Login: "mollen"}})

		app.EXPECT().GetCommentsByArticleID(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetCommentsByArticleID(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})
}


