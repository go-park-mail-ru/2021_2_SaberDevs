package models

import "context"

//Представление лайка
type LikeData struct {
	Ltype int `json:"type"`
	Sign  int `json:"sign"`
	Id    int `json:"id"`
}

type LikeDb struct {
	Id        int `json:"id"`
	ArticleId int `json:"articleid"`
	Signum    int `json:"sign"`
}

type GenericResponse struct {
	Status uint   `json:"status"`
	Data   string `json:"data"`
}

type LikesUsecase interface {
	Rating(ctx context.Context, a *LikeData, cValue string) (int, error)
}

type LikesRepository interface {
	Like(ctx context.Context, a *LikeData) (int, error)
	Dislike(ctx context.Context, a *LikeData) (int, error)
	Cancel(ctx context.Context, id int64) (int, error)
}
