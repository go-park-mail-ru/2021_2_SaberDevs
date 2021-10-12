package models

//Представление записи
type NewsRecord struct {
	Id           string   `json:"id"`
	PreviewUrl   string   `json:"previewUrl"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
	Text         string   `json:"text"`
	AuthorUrl    string   `json:"authorUrl"`
	AuthorName   string   `json:"authorName"`
	AuthorAvatar string   `json:"authorAvatar"`
	CommentsUrl  string   `json:"commentsUrl"`
	Comments     uint     `json:"comments"`
	Likes        uint     `json:"likes"`
}

//Тело ответа на API-call /getfeed

type RequestChunk struct {
	idLastLoaded string
	login        string
}

type ChunkResponse struct {
	Status    uint         `json:"status"`
	ChunkData []NewsRecord `json:"data"`
}
