create table if not exists article_likes(
		Id           serial primary key not null,
		articleId    int  references articles(Id) on delete cascade,
        signum       int 
		);