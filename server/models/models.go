package models

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Score    uint   `json:"score"`
}

type RequestUser struct {
	Login    string `json:"login"`
	// Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginBody struct {
	Login   string `json:"login"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type LoginResponse struct {
	Status uint      `json:"status"`
	Data   LoginBody `json:"data"`
	Msg    string    `json:"msg"`
}

type LogoutResponse struct {
	Status     uint   `json:"status"`
	GoodbyeMsg string `json:"goodbye"`
}

type SignUpBody struct {
	Login   string `json:"login"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Score   int    `json:"score"`
}

type RequestSignup struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type SignupResponse struct {
	Status uint       `json:"status"`
	SBody  SignUpBody `json:"data"`
	Msg    string     `json:"msg"`
}

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

type ErrorBody struct {
	Status   uint   `json:"status"`
	ErrorMsg string `json:"msg"`
}
