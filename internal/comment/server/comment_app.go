package main

import (
	"context"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
)

type CommentManager struct {
	usecase cmodels.CommentUsecase
}

func NewCommentManager(handler cmodels.CommentUsecase) *CommentManager {
	return &CommentManager{
		usecase: handler,
	}
}

func (m *CommentManager) CreateComment(ctx context.Context, comment *app.CreateCommentInput) (*app.CommentResponse, error) {
	modelComment := &cmodels.Comment{
		Id:          comment.Comment.Id,
		DateTime:    comment.Comment.DateTime,
		Text:        comment.Comment.Text,
		AuthorLogin: comment.Comment.AuthorLogin,
		ArticleId:   comment.Comment.ArticleId,
		ParentId:    comment.Comment.ParentId,
		IsEdited:    false,
		Likes:       int(comment.Comment.Likes),
	}
	response, err := m.usecase.CreateComment(ctx, modelComment, comment.SessionID)

	cmnt, _ := response.Data.(cmodels.PreparedComment)

	return &app.CommentResponse{
		Status: uint32(response.Status),
		Data: &app.PreparedComment{
			Id:        cmnt.Id,
			DateTime:  cmnt.DateTime,
			Text:      cmnt.Text,
			ArticleId: cmnt.ArticleId,
			ParentId:  cmnt.ParentId,
			IsEdited:  cmnt.IsEdited,
			Author:    &app.Author{
				Login:                cmnt.Author.Login,
				LastName:             cmnt.Author.Surname,
				FirstName:            cmnt.Author.Name,
				Score:                int32(cmnt.Author.Score),
				AvatarUrl:            cmnt.Author.AvatarURL,
			},
		},
		Msg: response.Msg,
	}, err
}

func (m *CommentManager) UpdateComment(ctx context.Context, comment *app.UpdateCommentInput) (*app.CommentResponse, error) {
	modelComment := &cmodels.Comment{
		Id:          comment.Comment.Id,
		DateTime:    comment.Comment.DateTime,
		Text:        comment.Comment.Text,
		AuthorLogin: comment.Comment.AuthorLogin,
		ArticleId:   comment.Comment.ArticleId,
		ParentId:    comment.Comment.ParentId,
		IsEdited:    false,
		Likes:       int(comment.Comment.Likes),
	}
	response, err := m.usecase.CreateComment(ctx, modelComment, comment.SessionID)

	cmnt, _ := response.Data.(cmodels.PreparedComment)

	return &app.CommentResponse{
		Status: uint32(response.Status),
		Data: &app.PreparedComment{
			Id:        cmnt.Id,
			DateTime:  cmnt.DateTime,
			Text:      cmnt.Text,
			ArticleId: cmnt.ArticleId,
			ParentId:  cmnt.ParentId,
			IsEdited:  cmnt.IsEdited,
			Author:    &app.Author{
				Login:                cmnt.Author.Login,
				LastName:             cmnt.Author.Surname,
				FirstName:            cmnt.Author.Name,
				Score:                int32(cmnt.Author.Score),
				AvatarUrl:            cmnt.Author.AvatarURL,
			},
		},
		Msg: response.Msg,
	}, err
}

func (m *CommentManager) GetCommentsByArticleID(ctx context.Context, articleID *app.ArticleID) (*app.CommentChunkResponse, error) {
	response, err := m.usecase.GetCommentsByArticleID(ctx, articleID.ArticleID)

	var preparedComment []*app.PreparedComment

	convertedResponse := response.Data.([]cmodels.PreparedComment)

	for _, c := range convertedResponse {
		preparedComment = append(preparedComment, &app.PreparedComment{
			Id:                   c.Id,
			DateTime:             c.DateTime,
			Text:                 c.Text,
			ArticleId:            c.ArticleId,
			ParentId:             c.ParentId,
			IsEdited:             c.IsEdited,
			Author:               &app.Author{
				Login:                c.Author.Login,
				LastName:             c.Author.Surname,
				FirstName:            c.Author.Name,
				Score:                int32(c.Author.Score),
				AvatarUrl:            c.Author.AvatarURL,
			},
			Likes:                int32(c.Likes),
		})
	}

	return &app.CommentChunkResponse{
		Status:               uint32(response.Status),
		Data:                 preparedComment,
		Msg:                  response.Msg,
	}, err
}
