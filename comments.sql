CREATE TABLE if not exists comments(
		Id          SERIAL PRIMARY KEY NOT NULL,
		AuthorLogin       VARCHAR(45) REFERENCES author(Login),
        ArticleId INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL ,
        ParentId INT REFERENCES comments(id) ON DELETE CASCADE,
		Text         TEXT,
		IsEdited bool,
		DateTime     VARCHAR(45)
		);
