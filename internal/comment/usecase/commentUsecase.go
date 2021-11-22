package usecase

import (
	"context"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
)

type commentUsecase struct {
	userRepo    umodels.UserRepository
	sessionRepo smodels.SessionRepository
	commentRepo cmodels.CommentRepository
}

func NewUserUsecase(ur umodels.UserRepository, sr smodels.SessionRepository, cr cmodels.CommentRepository) cmodels.CommentUsecase {
	return &commentUsecase{
		userRepo:    ur,
		sessionRepo: sr,
		commentRepo: cr,
	}
}

func (cu *commentUsecase) CreateComment(ctx context.Context, comment *cmodels.Comment, sessionID string) (cmodels.Response, error) {

}

func (cu *commentUsecase) UpdateComment(ctx context.Context, comment *cmodels.Comment, sessionID string) (cmodels.Response, error) {

}

func (cu *commentUsecase) GetCommentsByArticleID(ctx context.Context, articleID string) (cmodels.Response, error) {

}
