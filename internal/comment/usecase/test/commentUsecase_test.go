package test

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	amocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	cmock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models/mock"
	pnmocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models/mock"
	smocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	umocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models/mock"
	// umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	ucase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/usecase"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	pnRepo := pnmocks.NewMockPushNotificationRepository(ctrl)
	crepo := cmock.NewMockCommentRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().GetSessionLogin(gomock.Any(), gomock.Any()).Return("author", nil).AnyTimes()

		comment := cmodels.Comment{}
		crepo.EXPECT().StoreComment(gomock.Any(), gomock.Any()).Return(comment, nil).AnyTimes()

		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		article := amodels.FullArticle{}
		aRepo.EXPECT().GetByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(article, nil).AnyTimes()

		pnRepo.EXPECT().GetSubscription(gomock.Any(), gomock.Any()).Return(webpush.Subscription{}, errors.New("err")).AnyTimes()

		u := ucase.NewCommentUsecase(uRepo, sesRepo, crepo, pnRepo, aRepo)
		_, err := u.CreateComment(context.TODO(), &cmodels.Comment{}, "")
		assert.NoError(t, err)
	})
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	pnRepo := pnmocks.NewMockPushNotificationRepository(ctrl)
	crepo := cmock.NewMockCommentRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().GetSessionLogin(gomock.Any(), gomock.Any()).Return("author", nil).AnyTimes()

		comment := cmodels.Comment{AuthorLogin: "author"}
		crepo.EXPECT().GetCommentByID(gomock.Any(), gomock.Any()).Return(comment, nil).AnyTimes()

		crepo.EXPECT().UpdateComment(gomock.Any(), gomock.Any()).Return(comment, nil).AnyTimes()

		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		u := ucase.NewCommentUsecase(uRepo, sesRepo, crepo, pnRepo, aRepo)
		_, err := u.UpdateComment(context.TODO(), &cmodels.Comment{}, "")
		assert.NoError(t, err)
	})
}

func TestGetCommentsByArticleID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	pnRepo := pnmocks.NewMockPushNotificationRepository(ctrl)
	crepo := cmock.NewMockCommentRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		crepo.EXPECT().GetCommentsByArticleID(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()

		u := ucase.NewCommentUsecase(uRepo, sesRepo, crepo, pnRepo, aRepo)
		_, err := u.GetCommentsByArticleID(context.TODO(), 0)
		assert.NoError(t, err)
	})
}
