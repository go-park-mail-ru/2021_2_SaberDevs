package commentStream

import (
	"time"

	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
)

type repoChecker struct {
	pub         *Publisher
	commentRepo cmodels.CommentRepository
}

//easyjson:json
type sqlPreparedComment struct {
	Id          int64  `json:"Id"  db:"id"`
	Text        string `json:"text" db:"text"`
	ArticleId   int64  `json:"articleId" db:"articleid"`
	ArticleName string `json:"articleName" db:"title"`
	author      `json:"author"`
}

//easyjson:json
type author struct {
	Login     string `json:"login" db:"login"`
	Surname   string `json:"lastName" db:"surname"`
	Name      string `json:"firstName" db:"name"`
	AvatarURL string `json:"avatarUrl" db:"avatarurl"`
}

const checkWait = 15 * time.Second

func NewRepoChecker(p *Publisher, cr cmodels.CommentRepository) *repoChecker {
	return &repoChecker{
		pub:         p,
		commentRepo: cr,
	}
}

func (check *repoChecker) Run() {
	ticker := time.NewTicker(checkWait)
	// var lastId int64 = 0

	for {
		select {
		case <-ticker.C:
			comments, err := check.commentRepo.GetCommentsStream(0)
			if err != nil || len(comments) == 0 {
				continue
			}
			// if lastId < comments[0].Id {
			// 	lastId = comments[0].Id
			// }
			check.pub.broadcast <- comments
		}
	}
}
