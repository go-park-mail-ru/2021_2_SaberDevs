package usecase

import (
	"context"
	"net/http"
	"time"

	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/pkg/errors"
)

type commentUsecase struct {
	userRepo    umodels.UserRepository
	sessionRepo smodels.SessionRepository
	commentRepo cmodels.CommentRepository
}

func NewCommentUsecase(ur umodels.UserRepository, sr smodels.SessionRepository, cr cmodels.CommentRepository) cmodels.CommentUsecase {
	return &commentUsecase{
		userRepo:    ur,
		sessionRepo: sr,
		commentRepo: cr,
	}
}

func (cu *commentUsecase) CreateComment(ctx context.Context, comment *cmodels.Comment, sessionID string) (cmodels.Response, error) {
	login, err := cu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/CreateComment")
	}

	commentToStore := cmodels.Comment{
		DateTime:    time.Now().Format("2006/1/2 15:04"),
		Text:        comment.Text,
		AuthorLogin: login,
		ArticleId:   comment.ArticleId,
		Likes:       comment.Likes,
		ParentId:    comment.ParentId,
		IsEdited:    false,
	}

	storedComment, err := cu.commentRepo.StoreComment(ctx, &commentToStore)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/CreateComment")
	}

	userInRepo, err := cu.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/CreateComment")
	}

	responseData := cmodels.PreparedComment{
		Id:        storedComment.Id,
		DateTime:  storedComment.DateTime,
		Text:      storedComment.Text,
		ArticleId: storedComment.ArticleId,
		Likes:     storedComment.Likes,
		ParentId:  storedComment.ParentId,
		IsEdited:  storedComment.IsEdited,
		Author: cmodels.Author{
			Login:     userInRepo.Login,
			Surname:   userInRepo.Surname,
			Name:      userInRepo.Name,
			Score:     userInRepo.Score,
			AvatarURL: userInRepo.AvatarURL,
		},
	}
	response := cmodels.Response{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, nil
}

func (cu *commentUsecase) UpdateComment(ctx context.Context, comment *cmodels.Comment, sessionID string) (cmodels.Response, error) {
	login, err := cu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/UpdateComment")
	}

	commentInRepo, err := cu.commentRepo.GetCommentByID(ctx, comment.Id)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/UpdateComment")
	}
	if commentInRepo.AuthorLogin != login {
		return cmodels.Response{}, sbErr.ErrUnauthorized{
			Reason:   "request login doesnt match repo login",
			Function: "commentUsecase/UpdateComment",
		}
	}

	commentToUpdate := cmodels.Comment{
		Id:          comment.Id,
		Text:        comment.Text,
		Likes:       comment.Likes,
		AuthorLogin: login,
		IsEdited:    true,
	}

	updatedComment, err := cu.commentRepo.UpdateComment(ctx, &commentToUpdate)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/UpdateComment")
	}

	userInRepo, err := cu.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/UpdateComment")
	}

	responseData := cmodels.PreparedComment{
		Id:        updatedComment.Id,
		DateTime:  updatedComment.DateTime,
		Text:      updatedComment.Text,
		ArticleId: updatedComment.ArticleId,
		Likes:     updatedComment.Likes,
		ParentId:  updatedComment.ParentId,
		IsEdited:  updatedComment.IsEdited,
		Author: cmodels.Author{
			Login:     userInRepo.Login,
			Surname:   userInRepo.Surname,
			Name:      userInRepo.Name,
			Score:     userInRepo.Score,
			AvatarURL: userInRepo.AvatarURL,
		},
	}
	response := cmodels.Response{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, nil
}

func (cu *commentUsecase) GetCommentsByArticleID(ctx context.Context, articleID int64) (cmodels.Response, error) {
	responseData, err := cu.commentRepo.GetCommentsByArticleID(ctx, articleID, 0)
	if err != nil {
		return cmodels.Response{}, errors.Wrap(err, "commentUsecase/GetCommentsByArticleID")
	}

	response := cmodels.Response{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, nil
}
