package main

import (
	"context"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models/mock"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockCommentUsecase(ctrl)
	m := NewCommentManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.Response{}
		usecase.EXPECT().CreateComment(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()

		_, err := m.CreateComment(context.TODO(), &app.CreateCommentInput{Comment: &app.Comment{Likes: 9}})
		assert.NoError(t, err)
	})
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockCommentUsecase(ctrl)
	m := NewCommentManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.Response{}
		usecase.EXPECT().UpdateComment(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()

		_, err := m.UpdateComment(context.TODO(), &app.UpdateCommentInput{Comment: &app.Comment{Likes: 9}})
		assert.NoError(t, err)
	})
}

func TestGetCommentsByArticleID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockCommentUsecase(ctrl)
	m := NewCommentManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.Response{Data: []models.PreparedComment{}}
		usecase.EXPECT().GetCommentsByArticleID(gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()

		_, err := m.GetCommentsByArticleID(context.TODO(), &app.ArticleID{ArticleID: 9})
		assert.NoError(t, err)
	})
}
