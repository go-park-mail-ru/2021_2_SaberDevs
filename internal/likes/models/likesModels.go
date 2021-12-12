package models

import "context"

//Представление лайка
type LikeData struct {
	ltype int `json:"type"`
	sign  int `json:"sign"`
	id    int `json:"id"`
}

type LikeDb struct {
	ltype  int `json:"type"`
	sign   int `json:"sign"`
	id     int `json:"id"`
	userId int `json:"id"`
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
