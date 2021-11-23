drop table if exists comments;

CREATE TABLE if not exists comments(
		Id          SERIAL PRIMARY KEY NOT NULL,
		AuthorLogin       VARCHAR(45) REFERENCES author(Login),
        ArticleId INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL ,
        ParentId INT REFERENCES comments(id) ON DELETE CASCADE,
		Text         TEXT,
		IsEdited bool,
		DateTime     VARCHAR(45)
		);

insert into comments (AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime) values ('mollenTEST1', 1, null, 'крутой комент1', false, '2021/11/23 13:13');
insert into comments (AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime) values ('mollenTEST1', 1, null, 'крутой комент2', false, '2021/11/23 13:13');
insert into comments (AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime) values ('mollenTEST1', 1, null, 'крутой комент3', false, '2021/11/23 13:13');
